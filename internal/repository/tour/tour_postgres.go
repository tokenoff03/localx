package tour

import (
	"localx/internal/models"

	"github.com/jmoiron/sqlx"
)

type TourPostgres struct {
	db *sqlx.DB
}

func NewTourPostgres(db *sqlx.DB) *TourPostgres {
	return &TourPostgres{db: db}
}

func (t *TourPostgres) CreateTour(tour models.Tour, companyId int) (int, error) {

	return 0, nil
}

func (t *TourPostgres) GetById(id int) (models.Tour, error) {

	return models.Tour{}, nil
}
