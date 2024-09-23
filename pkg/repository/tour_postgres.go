package repository

import (
	"localx"

	"github.com/jmoiron/sqlx"
)

type TourPostgres struct {
	db *sqlx.DB
}

func NewTourPostgres(db *sqlx.DB) *TourPostgres {
	return &TourPostgres{db: db}
}

func (t *TourPostgres) CreateTour(tour localx.Tour, companyId int) (int, error) {

	return 0, nil
}

func (t *TourPostgres) GetById(id int) (localx.Tour, error) {

	return localx.Tour{}, nil
}
