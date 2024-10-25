package tour

import (
	"context"
	"database/sql"
	"localx/internal/models"
	"time"

	"github.com/jmoiron/sqlx"
)

type TourPostgres struct {
	db *sqlx.DB
}

func NewTourPostgres(db *sqlx.DB) *TourPostgres {
	return &TourPostgres{db: db}
}

func (t *TourPostgres) CreateTour(ctx context.Context, tour models.Tour, companyId int) (int, error) {
	query := `
	INSERT INTO tour (id, company_id, title, start_time, end_time, group_size, languages, free_cancellation, cancellation_condition, description, meeting_place, 
	arrival_place, what_is_included, what_to_prepare, prohibitions, price, images)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9::json, $10, $11, $12, $13, $14, $15, $16, $17) RETURNING id`
	var id int
	row := t.db.QueryRowContext(ctx, query, tour.ID, companyId, tour.Title, tour.StartTime, tour.EndTime, tour.GroupSize, tour.Languages, tour.FreeCancellation,
		tour.CancellationCondition, tour.Description, tour.MeetingPlace, tour.ArrivalPlace, tour.WhatIsIncluded, tour.WhatToPrepare, tour.Prohibitions,
		tour.Price, tour.Images)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return int(id), nil
}

func (t *TourPostgres) GetTourById(ctx context.Context, id int) (models.Tour, error) {
	var tour models.Tour
	query := `
	SELECT id, company_id, title, start_time, end_time, group_size, languages, free_cancellation, cancellation_condition, description, meeting_place, 
	arrival_place, what_is_included, what_to_prepare, prohibitions, price, images
	FROM tour
	WHERE id = $1
	`

	err := t.db.GetContext(ctx, &tour, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.Tour{}, nil
		}
		return models.Tour{}, err
	}

	return tour, nil
}

func (t *TourPostgres) UpdateTour(ctx context.Context, tour models.Tour) error {
	query := `
	UPDATE tour
	SET title = $1, start_time = $2, end_time = $3, group_size = $4, languages = $5, free_cancellation = $6, cancellation_condition = $7, description = $8, 
	meeting_place = $9, arrival_place = $10, what_is_included = $11, what_to_prepare = $12, prohibitions = $13, price = $14, images = $15
	WHERE id = $16
	`
	_, err := t.db.ExecContext(ctx, query, tour.Title, tour.StartTime, tour.EndTime, tour.GroupSize, tour.Languages, tour.FreeCancellation,
		tour.CancellationCondition, tour.Description, tour.MeetingPlace, tour.ArrivalPlace, tour.WhatIsIncluded, tour.WhatToPrepare,
		tour.Prohibitions, tour.Price, tour.Images, tour.ID,
	)
	return err
}

func (t *TourPostgres) DeleteTour(ctx context.Context, id int) error {
	query := "DELETE FROM tour WHERE id = $1"
	_, err := t.db.ExecContext(ctx, query, id)
	return err
}

func (t *TourPostgres) SetTitle(ctx context.Context, id int, title string) error {
	query := "UPDATE tour SET title = $1 WHERE id = $2"
	_, err := t.db.ExecContext(ctx, query, title, id)
	return err
}

func (t *TourPostgres) SetStartTime(ctx context.Context, id int, startTime time.Time) error {
	query := "UPDATE tour SET start_time = $1 WHERE id = $2"
	_, err := t.db.ExecContext(ctx, query, startTime, id)
	return err
}

func (t *TourPostgres) SetEndTime(ctx context.Context, id int, endTime time.Time) error {
	query := "UPDATE tour SET end_time = $1 WHERE id = $2"
	_, err := t.db.ExecContext(ctx, query, endTime, id)
	return err
}

func (t *TourPostgres) SetLanguages(ctx context.Context, id int, languages string) error {
	query := "UPDATE tour SET languages = $1 WHERE id = $2"
	_, err := t.db.ExecContext(ctx, query, languages, id)
	return err
}

func (t *TourPostgres) SetFreeCancellation(ctx context.Context, id int, freeCancellation bool) error {
	query := "UPDATE tour SET free_cancellation = $1 WHERE id = $2"
	_, err := t.db.ExecContext(ctx, query, freeCancellation, id)
	return err
}

func (t *TourPostgres) SetCancellationCondition(ctx context.Context, id int, cancellationCondition string) error {
	query := "UPDATE tour SET cancellation_condition = $1 WHERE id = $2"
	_, err := t.db.ExecContext(ctx, query, cancellationCondition, id)
	return err
}

func (t *TourPostgres) SetDescription(ctx context.Context, id int, description string) error {
	query := "UPDATE tour SET description = $1 WHERE id = $2"
	_, err := t.db.ExecContext(ctx, query, description, id)
	return err
}

func (t *TourPostgres) SetMeetingPlace(ctx context.Context, id int, meetingPlace string) error {
	query := "UPDATE tour SET meeting_place = $1 WHERE id = $2"
	_, err := t.db.ExecContext(ctx, query, meetingPlace, id)
	return err
}

func (t *TourPostgres) SetArrivalPlace(ctx context.Context, id int, arrivalPlace string) error {
	query := "UPDATE tour SET arrival_place = $1 WHERE id = $2"
	_, err := t.db.ExecContext(ctx, query, arrivalPlace, id)
	return err
}

func (t *TourPostgres) SetWhatIsIncluded(ctx context.Context, id int, whatIsIncluded string) error {
	query := "UPDATE tour SET what_is_included = $1 WHERE id = $2"
	_, err := t.db.ExecContext(ctx, query, whatIsIncluded, id)
	return err
}

func (t *TourPostgres) SetWhatToPrepare(ctx context.Context, id int, whatToPrepare string) error {
	query := "UPDATE tour SET what_to_prepare = $1 WHERE id = $2"
	_, err := t.db.ExecContext(ctx, query, whatToPrepare, id)
	return err
}

func (t *TourPostgres) SetProhibitions(ctx context.Context, id int, prohibitions string) error {
	query := "UPDATE tour SET prohibitions = $1 WHERE id = $2"
	_, err := t.db.ExecContext(ctx, query, prohibitions, id)
	return err
}

func (t *TourPostgres) SetPrice(ctx context.Context, id int, price int) error {
	query := "UPDATE tour SET price = $1 WHERE id = $2"
	_, err := t.db.ExecContext(ctx, query, price, id)
	return err
}

func (t *TourPostgres) SetImages(ctx context.Context, id int, images string) error {
	query := "UPDATE tour SET images = $1 WHERE id = $2"
	_, err := t.db.ExecContext(ctx, query, images, id)
	return err
}
