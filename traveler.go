package localx

import "time"

type Traveler struct {
	ID            int       `json:"id"`
	FullName      string    `json:"full_name"`
	Email         string    `json:"email"`
	PhoneNumber   string    `json:"phone_number"`
	Instagram     *string   `json:"instagram"` // NULLable field
	DateOfBirth   time.Time `json:"date_of_birth"`
	Password      string    `json:"password"`
	City          string    `json:"city"`
	Country       string    `json:"country"`
	Description   *string   `json:"description"` // NULLable field
	Interest      int       `json:"interest"`
	FavoriteTours int       `json:"favorite_tours"`
}

type Interests struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}
