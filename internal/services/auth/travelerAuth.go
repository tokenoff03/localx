package auth

import (
	"localx/internal/models"
	"localx/internal/repository"
	"sync"
	"time"
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
	accessTokens  map[string]string // access token по email
	refreshTokens map[string]string // refresh token по email
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
