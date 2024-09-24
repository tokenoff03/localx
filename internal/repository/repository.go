package repository

import (
	"localx/internal/models"
	"localx/internal/repository/tour"

	"github.com/jmoiron/sqlx"
)

type Tour interface {
	CreateTour(tour models.Tour, companyId int) (int, error)
	GetById(id int) (models.Tour, error)
}

type Repository struct {
	Tour
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{Tour: tour.NewTourPostgres(db)}
}
