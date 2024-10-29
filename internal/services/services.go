package services

import (
	"context"
	"localx/internal/models"
	"localx/internal/repository"
	"localx/internal/services/auth"
	"localx/internal/services/tour"
)

type Tour interface {
	CreateTour(ctx context.Context, tour models.Tour, companyId int64) (int64, error)
	GetTourById(ctx context.Context, id int64) (models.Tour, error)
	UpdateTour(ctx context.Context, tour models.Tour) error
	DeleteTour(ctx context.Context, id int64) error
	PartialUpdateTour(ctx context.Context, tour models.Tour) error
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
