package repository

import (
	"context"
	"database/sql"
	"fmt"
	"localx/internal/models"

	"github.com/jmoiron/sqlx"
)

type TravelerRepository struct {
	db *sqlx.DB
}

func NewTravelerRepository(db *sqlx.DB) *TravelerRepository {
	return &TravelerRepository{
		db: db,
	}
}

// CreateTraveler создает новую запись о путешественнике в базе данных.
func (r *TravelerRepository) CreateTraveler(ctx context.Context, traveler models.Traveler) (int64, error) {
	query := `
		INSERT INTO traveler (first_name, last_name, email, phone_number, instagram, date_of_birth, city, country, description, interest, favorite_tours, profile_pic)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12) RETURNING id`
	var id int64
	err := r.db.QueryRowContext(ctx, query, traveler.FirstName, traveler.LastName, traveler.Email, traveler.PhoneNumber, traveler.Instagram, traveler.DateOfBirth, traveler.City, traveler.Country, traveler.Description, traveler.Interest, traveler.FavoriteTours, traveler.ProfilePictureURL).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

// GetTravelerByID возвращает запись о путешественнике по его ID.
func (r *TravelerRepository) GetTravelerByID(ctx context.Context, id int64) (*models.Traveler, error) {
	var traveler models.Traveler
	query := `
		SELECT id, first_name, last_name, email, phone_number, instagram, date_of_birth, city, country, description, interest, favorite_tours, profile_pic
		FROM traveler
		WHERE id = $1`
	err := r.db.GetContext(ctx, &traveler, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &traveler, nil
}

// UpdateTraveler обновляет информацию о путешественнике в базе данных.
func (r *TravelerRepository) UpdateTraveler(ctx context.Context, traveler models.Traveler) error {
	query := `
		UPDATE traveler
		SET first_name = $1, last_name = $2, email = $3, phone_number = $4, instagram = $5, date_of_birth = $6, city = $7, country = $8, description = $9, interest = $10, favorite_tours = $11, profile_pic = $12
		WHERE id = $13`
	_, err := r.db.ExecContext(ctx, query, traveler.FirstName, traveler.LastName, traveler.Email, traveler.PhoneNumber, traveler.Instagram, traveler.DateOfBirth, traveler.City, traveler.Country, traveler.Description, traveler.Interest, traveler.FavoriteTours, traveler.ProfilePictureURL, traveler.ID)
	return err
}

// DeleteTraveler удаляет запись о путешественнике из базы данных.
func (r *TravelerRepository) DeleteTraveler(ctx context.Context, id int64) error {
	query := "DELETE FROM traveler WHERE id = $1"
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}

// SetProfilePicture обновляет ссылку на фотографию профиля путешественника.
func (r *TravelerRepository) SetProfilePicture(ctx context.Context, id int64, profilePicture string) error {
	query := "UPDATE traveler SET profile_pic = $1 WHERE id = $2"
	_, err := r.db.ExecContext(ctx, query, profilePicture, id)
	return err
}

// GetTravelerByEmail возвращает запись о путешественнике по его email.
func (r *TravelerRepository) GetTravelerByEmail(ctx context.Context, email string) (*models.Traveler, error) {
	var traveler models.Traveler
	query := `
		SELECT id, first_name, last_name, email, phone_number, instagram, date_of_birth, city, country, description, interest, favorite_tours, profile_pic
		FROM traveler
		WHERE email = $1`
	err := r.db.GetContext(ctx, &traveler, query, email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &traveler, nil
}

// Дополнительные методы для обновления отдельных полей

func (r *TravelerRepository) SetDescription(ctx context.Context, id int64, description string) error {
	query := "UPDATE traveler SET description = $1 WHERE id = $2"
	_, err := r.db.ExecContext(ctx, query, description, id)
	return err
}

func (r *TravelerRepository) SetCity(ctx context.Context, id int64, city string) error {
	query := "UPDATE traveler SET city = $1 WHERE id = $2"
	_, err := r.db.ExecContext(ctx, query, city, id)
	return err
}

func (r *TravelerRepository) SetInstagram(ctx context.Context, id int64, instagram string) error {
	query := "UPDATE traveler SET instagram = $1 WHERE id = $2"
	_, err := r.db.ExecContext(ctx, query, instagram, id)
	return err
}

func (r *TravelerRepository) SetPhoneNumber(ctx context.Context, id int64, phoneNumber string) error {
	query := "UPDATE traveler SET phone_number = $1 WHERE id = $2"
	_, err := r.db.ExecContext(ctx, query, phoneNumber, id)
	return err
}

func (r *TravelerRepository) SetInterest(ctx context.Context, id int64, interest int) error {
	query := "UPDATE traveler SET interest = $1 WHERE id = $2"
	_, err := r.db.ExecContext(ctx, query, interest, id)
	return err
}

func (r *TravelerRepository) ClearProfilePicture(ctx context.Context, userID int64) error {
	query := `UPDATE traveler SET profile_pic = NULL WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, userID)
	if err != nil {
		return fmt.Errorf("failed to clear profile picture: %w", err)
	}
	return nil
}
