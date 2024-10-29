package tour

import (
	"context"
	"database/sql"
	"localx/internal/models"

	"github.com/jmoiron/sqlx"
)

type TourPostgres struct {
	db *sqlx.DB
}

func NewTourPostgres(db *sqlx.DB) *TourPostgres {
	return &TourPostgres{db: db}
}

func (t *TourPostgres) CreateTour(ctx context.Context, tour models.Tour, companyId int64) (int64, error) {
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

	return int64(id), nil
}

func (t *TourPostgres) GetTourById(ctx context.Context, id int64) (models.Tour, error) {
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

func (t *TourPostgres) DeleteTour(ctx context.Context, id int64) error {
	query := "DELETE FROM tour WHERE id = $1"
	_, err := t.db.ExecContext(ctx, query, id)
	return err
}

func (t *TourPostgres) UpdateTourDetails(ctx context.Context, tour models.Tour) error {
	query := `
    UPDATE tour SET
    title = COALESCE($1, title),
    start_time = COALESCE($2, start_time),
    end_time = COALESCE($3, end_time),
    group_size = COALESCE($4, group_size),
    languages = COALESCE($5, languages),
    free_cancellation = $6,
    cancellation_condition = COALESCE($7, cancellation_condition),
    description = COALESCE($8, description),
    meeting_place = COALESCE($9, meeting_place),
    arrival_place = COALESCE($10, arrival_place),
    what_is_included = COALESCE($11, what_is_included),
    what_to_prepare = COALESCE($12, what_to_prepare),
    prohibitions = COALESCE($13, prohibitions),
    price = COALESCE($14, price),
    images = COALESCE($15, images)
    WHERE id = $16
    `

	_, err := t.db.ExecContext(ctx, query,
		tour.Title, tour.StartTime, tour.EndTime, tour.GroupSize, tour.Languages,
		tour.FreeCancellation, tour.CancellationCondition, tour.Description,
		tour.MeetingPlace, tour.ArrivalPlace, tour.WhatIsIncluded, tour.WhatToPrepare,
		tour.Prohibitions, tour.Price, tour.Images, tour.ID,
	)
	return err
}
