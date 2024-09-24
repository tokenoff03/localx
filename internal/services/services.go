package services

import (
	"localx/internal/models"
	"localx/internal/repository"
	"localx/internal/services/tour"
)

type Tour interface {
	CreateTour(tour models.Tour, companyId int) (int, error)
	GetById(id int) (models.Tour, error)
}

type Services struct {
	Tour
}

func NewServices(repo *repository.Repository) *Services {
	return &Services{Tour: tour.NewTour(repo)}
}
