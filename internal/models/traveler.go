package models

import "time"

type Traveler struct {
	ID            int       `json:"id" db:"id"`
	FirstName     string    `json:"first_name" db:"first_name"`
	LastName      string    `json:"last_name" db:"last_name"`
	Email         string    `json:"email" db:"email"`
	PhoneNumber   string    `json:"phone_number" db:"phone_number"`
	Instagram     *string   `json:"instagram" db:"instagram"` // NULLable field
	DateOfBirth   time.Time `json:"date_of_birth" db:"date_of_birth"`
	City          string    `json:"city" db:"city"`
	Country       string    `json:"country" db:"country"`
	Description   *string   `json:"description" db:"description"` // NULLable field
	Interest      int       `json:"interest" db:"interest"`
	FavoriteTours int       `json:"favorite_tours" db:"favorite_tours"`
}

type Interests struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}
