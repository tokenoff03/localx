package models

import (
	"encoding/json"
	"time"
)

type Tour struct {
	ID                    int             `db:"id" json:"id"`
	CompanyID             int             `db:"company_id" json:"company_id"`
	Title                 string          `db:"title" json:"title"`
	StartTime             time.Time       `db:"start_time" json:"start_time"`
	EndTime               time.Time       `db:"end_time" json:"end_time"`
	GroupSize             int             `db:"group_size" json:"group_size"`
	Languages             string          `db:"languages" json:"languages"`
	FreeCancellation      bool            `db:"free_cancellation" json:"free_cancellation"`
	CancellationCondition json.RawMessage `db:"cancellation_condition" json:"cancellation_condition"`
	Description           string          `db:"description" json:"description"`
	MeetingPlace          string          `db:"meeting_place" json:"meeting_place"`
	ArrivalPlace          string          `db:"arrival_place" json:"arrival_place"`
	WhatIsIncluded        string          `db:"what_is_included" json:"what_is_included"`
	WhatToPrepare         string          `db:"what_to_prepare" json:"what_to_prepare"`
	Prohibitions          string          `db:"prohibitions" json:"prohibitions"`
	Price                 int             `db:"price" json:"price"`
	Images                string          `db:"images" json:"images"`
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
