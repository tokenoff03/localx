package models

type Company struct {
	ID          int64   `json:"id"`
	Name        string  `json:"name"`
	PhoneNumber *string `json:"phone_number"` // NULLable field
	Instagram   *string `json:"instagram"`    // NULLable field
	Gis         *string `json:"2gis"`         // NULLable field
	Email       string  `json:"email"`
}

type CompanyReview struct {
	ID         int64   `json:"id"`
	Rating     float64 `json:"rating"`
	Text       string  `json:"text"`
	TravelerID int64   `json:"traveler_id"`
	CompanyID  int64   `json:"company_id"`
}
