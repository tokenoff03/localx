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

func (t *AuthTraveler) GetTravelerById(id int) (models.Traveler, error) {
	var traveler models.Traveler

	query := fmt.Sprintf("SELECT t.id, t.first_name, t.last_name, t.email, t.phone_number, t.instagram, t.date_of_birth, t.city, t.country, t.description, t.interest, t.favorite_tours FROM %s t WHERE t.id=$1", "traveler")

	err := t.db.Get(&traveler, query, id)

	if err != nil {
		return traveler, nil
	}

	return traveler, err
}
func (t *AuthTraveler) GetTravelerByEmail(email string) (models.Traveler, error) {
	var traveler models.Traveler

	query := fmt.Sprintf("SELECT t.id, t.first_name, t.last_name, t.email, t.phone_number, t.instagram, t.date_of_birth, t.city, t.country, t.description, t.interest, t.favorite_tours FROM %s t WHERE t.email=$1", "traveler")

	err := t.db.Get(&traveler, query, email)

	if errors.Is(err, sql.ErrNoRows) {

		return traveler, nil
	}

	return traveler, err
}

func (t *AuthTraveler) CreateTraveler(traveler models.TravelerSignUp) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (first_name,last_name,email, phone_number, date_of_birth, city, country, interest) values ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id", "traveler")
	row := t.db.QueryRow(query, traveler.FirstName, traveler.LastName, traveler.Email, traveler.PhoneNumber, traveler.DateOfBirth, traveler.City, traveler.Country, traveler.Interest)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil

}
