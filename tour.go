package localx

import "time"

type Tour struct {
	ID                    int       `json:"id"`
	CompanyID             int       `json:"company_id"`
	Title                 string    `json:"title"`
	StartTime             time.Time `json:"start_time"`
	EndTime               time.Time `json:"end_time"`
	GroupSize             int       `json:"group_size"`
	Languages             string    `json:"languages"`
	FreeCancellation      bool      `json:"free_cancellation"`
	CancellationCondition string    `json:"cancellation_condition"` // JSON type
	Description           string    `json:"description"`
	MeetingPlace          string    `json:"meeting_place"`
	ArrivalPlace          string    `json:"arrival_place"`
	WhatIsIncluded        string    `json:"what_is_included"`
	WhatToPrepare         string    `json:"what_to_prepare"`
	Prohibitions          string    `json:"prohibitions"`
	Price                 int       `json:"price"`
	Images                string    `json:"images"`
}

type TourModerator struct {
	ID         int `json:"id"`
	EmployeeID int `json:"employee_id"`
	TourID     int `json:"tour_id"`
}

type TourReview struct {
	ID         int     `json:"id"`
	Rating     float64 `json:"rating"`
	Text       string  `json:"text"`
	TourID     int     `json:"tour_id"`
	TravelerID int     `json:"traveler_id"`
}

type TourGroup struct {
	ID         int `json:"id"`
	TourID     int `json:"tour_id"`
	TravelerID int `json:"traveler_id"`
}
