package repository

import (
	"github.com/jmoiron/sqlx"
)

func NewPostgresDB(urlPath string) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", urlPath)

	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
