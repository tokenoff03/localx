package auth

import (
	"database/sql"
	"errors"
	"fmt"
	"localx/internal/models"

	"github.com/jmoiron/sqlx"
)

type AuthTraveler struct {
	db *sqlx.DB
}

func NewAuthTavelerPostgres(db *sqlx.DB) *AuthTraveler {
	return &AuthTraveler{db: db}
}

func (t *AuthTraveler) GetAllTraveler() ([]models.Traveler, error) {
	var travelers []models.Traveler

	query := fmt.Sprintf("SELECT t.id, t.first_name, t.last_name, t.email, t.phone_number, t.instagram, t.date_of_birth, t.city, t.country, t.description, t.interest, t.favorite_tours FROM %s t", "traveler")
	err := t.db.Select(&travelers, query)
	if err != nil {
		return nil, err
	}

	return travelers, nil
}
func (t *AuthTraveler) GetTraveler(email string) (models.Traveler, error) {
	var traveler models.Traveler

	query := fmt.Sprintf("SELECT t.id, t.first_name, t.last_name, t.email, t.phone_number, t.instagram, t.date_of_birth, t.city, t.country, t.description, t.interest, t.favorite_tours FROM %s t WHERE t.email=$1", "traveler")

	err := t.db.Get(&traveler, query, email)

	if errors.Is(err, sql.ErrNoRows) {

		return traveler, nil
	}

	return traveler, err
}

func (t *AuthTraveler) CreateTraveler(traveler models.Traveler) (int, error) {
	return 0, nil
}
