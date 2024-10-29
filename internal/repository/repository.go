package repository

import (
	"context"
	"localx/internal/models"
	"localx/internal/repository/auth"
	"localx/internal/repository/tour"

	"github.com/jmoiron/sqlx"
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
	GetTravelerByEmail(email string) (models.Traveler, error)
	GetAllTraveler() ([]models.Traveler, error)
	GetTravelerById(id int) (models.Traveler, error)
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
