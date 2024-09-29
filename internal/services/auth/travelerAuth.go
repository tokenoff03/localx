package auth

import (
	"errors"
	"localx/internal/models"
	"localx/internal/repository"
	"sync"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	tokenTTL        = 15 * time.Minute //15 minutes
	signingKey      = "213easdxz1c856eq"
	refreshTokenTTL = 7 * 24 * time.Hour //7 days
)

type otpData struct {
	code      string
	expiresAt time.Time
}

type InMemoryTokenStorage struct {
	accessTokens  map[string]string // access token по phoneNumber
	refreshTokens map[string]string // refresh token по phoneNumber
}

type AuthTravelerService struct {
	repo     *repository.Repository
	otpStore map[string]otpData
	tokens   *InMemoryTokenStorage
	mu       sync.Mutex
}

func NewAuthTravelerService(rep *repository.Repository) *AuthTravelerService {
	return &AuthTravelerService{
		repo:     rep,
		otpStore: make(map[string]otpData),
		tokens: &InMemoryTokenStorage{
			accessTokens:  make(map[string]string),
			refreshTokens: make(map[string]string),
		},
	}
}

func (t *AuthTravelerService) CreateTraveler(traveler models.TravelerSignUp) (int, error) {
	return t.repo.AuthTraveler.CreateTraveler(traveler)
}

func (t *AuthTravelerService) GetTravelerById(id int) (models.Traveler, error) {
	return t.repo.AuthTraveler.GetTravelerById(id)
}

func (t *AuthTravelerService) GetTravelerByEmail(email string) (models.Traveler, error) {
	return t.repo.AuthTraveler.GetTravelerByEmail(email)
}

func (t *AuthTravelerService) GetAllTraveler() ([]models.Traveler, error) {
	return t.repo.AuthTraveler.GetAllTraveler()
}

func (t *AuthTravelerService) GenerateToken(id int) (string, error) {
	traveler, err := t.GetTravelerById(id)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(tokenTTL).Unix(),
		Subject:   string(traveler.ID),
	})
	return token.SignedString([]byte(signingKey))
}

func (s *AuthTravelerService) ParseToken(accessToken string) (string, error) {
	token, err := jwt.Parse(accessToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(signingKey), nil
	})
	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("token claims are not of type *tokenClaims")
	}

	return claims["sub"].(string), nil
}

func (t *AuthTravelerService) GenerateRefreshToken(id int) (string, error) {
	traveler, err := t.GetTravelerById(id)
	if err != nil {
		return "", err
	}

	// Создаем refresh токен с длительным сроком действия (например, 7 дней)
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(refreshTokenTTL).Unix(),
		Subject:   string(traveler.ID),
	})
	return refreshToken.SignedString([]byte(signingKey))
}
