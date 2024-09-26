package user

import (
	"context"
	"database/sql"
	"localx/internal/models"

	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (u UserRepository) CreateUser(ctx context.Context, user models.User) (int64, error) {
	query := `
		INSERT INTO users (name, profile_picture, description, email, phone, city, instagram, interests)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id`
	var id int64
	err := u.db.QueryRowContext(ctx, query, user.Name, user.ProfilePicture, user.Description, user.Email, user.Phone, user.City, user.Instagram, user.Interests).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (u UserRepository) GetUserByID(ctx context.Context, id int64) (*models.User, error) {
	var user models.User
	query := `
		SELECT id, name, profile_picture, description, email, phone, city, instagram, interests
		FROM users
		WHERE id = $1`
	err := u.db.GetContext(ctx, &user, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (u UserRepository) UpdateUser(ctx context.Context, user models.User) error {
	query := `
		UPDATE users
		SET name = $1, profile_picture = $2, description = $3, email = $4, phone = $5, city = $6, instagram = $7, interests = $8
		WHERE id = $9`
	_, err := u.db.ExecContext(ctx, query, user.Name, user.ProfilePicture, user.Description, user.Email, user.Phone, user.City, user.Instagram, user.Interests, user.ID)
	return err
}

func (u UserRepository) DeleteUser(ctx context.Context, id int64) error {
	query := "DELETE FROM users WHERE id = $1"
	_, err := u.db.ExecContext(ctx, query, id)
	return err
}

func (u UserRepository) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	query := `
		SELECT id, name, profile_picture, description, email, phone, city, instagram, interests
		FROM users
		WHERE email = $1`
	err := u.db.GetContext(ctx, &user, query, email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (u UserRepository) SetDescription(ctx context.Context, id int64, description string) error {
	query := "UPDATE users SET description = $1 WHERE id = $2"
	_, err := u.db.ExecContext(ctx, query, description, id)
	return err
}

func (u UserRepository) SetCity(ctx context.Context, id int64, city string) error {
	query := "UPDATE users SET city = $1 WHERE id = $2"
	_, err := u.db.ExecContext(ctx, query, city, id)
	return err
}

func (u UserRepository) SetInstagram(ctx context.Context, id int64, instagram string) error {
	query := "UPDATE users SET instagram = $1 WHERE id = $2"
	_, err := u.db.ExecContext(ctx, query, instagram, id)
	return err
}

func (u UserRepository) SetProfilePicture(ctx context.Context, id int64, profilePicture string) error {
	query := "UPDATE users SET profile_picture = $1 WHERE id = $2"
	_, err := u.db.ExecContext(ctx, query, profilePicture, id)
	return err
}

func (u UserRepository) SetPhone(ctx context.Context, id int64, phone string) error {
	query := "UPDATE users SET phone = $1 WHERE id = $2"
	_, err := u.db.ExecContext(ctx, query, phone, id)
	return err
}

func (u UserRepository) SetInterests(ctx context.Context, id int64, interests string) error {
	query := "UPDATE users SET interests = $1 WHERE id = $2"
	_, err := u.db.ExecContext(ctx, query, interests, id)
	return err
}
