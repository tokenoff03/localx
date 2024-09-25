package auth

import (
	"errors"
	"localx/internal/models"
	"localx/internal/repository"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	salt       = "asdasdqw16531865zxcq1"
	tokenTTL   = 15 * time.Minute
	signingKey = "213easdxz1c856eq"
)

type OTPStorage struct {
	code string
}

type AuthTravelerService struct {
	repo *repository.Repository
}

func NewAuthTravelerService(rep *repository.Repository) *AuthTravelerService {
	return &AuthTravelerService{repo: rep}
}

func (t *AuthTravelerService) CreateTraveler(traveler models.Traveler) (int, error) {
	return 0, nil
}

func (t *AuthTravelerService) GetTraveler(phoneNumber string) (models.Traveler, error) {
	return t.repo.AuthTraveler.GetTraveler(phoneNumber)
}

func (t *AuthTravelerService) GetAllTraveler() ([]models.Traveler, error) {
	return t.repo.AuthTraveler.GetAllTraveler()
}

func (t *AuthTravelerService) GenerateToken(phoneNumber string) (string, error) {
	traveler, err := t.GetTraveler(phoneNumber)
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
