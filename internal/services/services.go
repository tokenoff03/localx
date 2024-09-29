package services

import (
	"localx/internal/models"
	"localx/internal/repository"
	"localx/internal/services/auth"
	"localx/internal/services/tour"
)

type Tour interface {
	CreateTour(tour models.Tour, companyId int) (int, error)
	GetById(id int) (models.Tour, error)
}

type AuthTraveler interface {
	CreateTraveler(traveler models.TravelerSignUp) (int, error)
	GetTravelerById(id int) (models.Traveler, error)
	GetTravelerByEmail(email string) (models.Traveler, error)
	GenerateToken(id int) (string, error)
	ParseToken(accessToken string) (string, error)
	GetAllTraveler() ([]models.Traveler, error)
	GenerateRefreshToken(id int) (string, error)
	GenerateAndSaveOTP(email string) (string, error)
	ValidateOTP(email, inputCode string) (bool, error)
	CleanExpiredOTPs()
	SendEmail(to string, subject string, body string) error
	GetTokens(phoneNumber string) (string, string, error)
	StoreTokens(phoneNumber, accessToken, refreshToken string)
	UpdateTokens(phoneNumber, accessToken, refreshToken string) error
}

type Services struct {
	Tour
	AuthTraveler
}

func NewServices(repo *repository.Repository) *Services {
	return &Services{
		Tour:         tour.NewTour(repo),
		AuthTraveler: auth.NewAuthTravelerService(repo),
	}
}
