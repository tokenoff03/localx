package repository

import (
	"localx/internal/models"
	"localx/internal/repository/auth"
	"localx/internal/repository/tour"

	"github.com/jmoiron/sqlx"
)

type Tour interface {
	CreateTour(tour models.Tour, companyId int) (int, error)
	GetById(id int) (models.Tour, error)
}

type AuthTraveler interface {
	CreateTraveler(traveler models.Traveler) (int, error)
	GetTraveler(phoneNumber string) (models.Traveler, error)
	GetAllTraveler() ([]models.Traveler, error)
}

type Repository struct {
	Tour
	AuthTraveler
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Tour:         tour.NewTourPostgres(db),
		AuthTraveler: auth.NewAuthTavelerPostgres(db),
	}
}
