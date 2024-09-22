package repository

import "github.com/jmoiron/sqlx"

type Tour interface {
	CreateTour(tour localx.Tour, companyId int) (int, error)
	GetById(id int) (localx.Tour, error)
}

type Repository struct {
	Tour
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{Tour: }
}
