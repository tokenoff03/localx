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
	CreateTraveler(traveler models.Traveler) (int, error)
	GetTraveler(phoneNumber string) (models.Traveler, error)
	GenerateToken(phoneNumber string) (string, error)
	ParseToken(accessToken string) (string, error)
	GetAllTraveler() ([]models.Traveler, error)
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
