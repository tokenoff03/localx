package repository

import (
	"context"
	"localx/internal/models"
	"localx/internal/repository/auth"
	"localx/internal/repository/tour"
	"time"

	"github.com/jmoiron/sqlx"
)

type Tour interface {
	CreateTour(ctx context.Context, tour models.Tour, companyId int) (int, error)
	GetTourById(ctx context.Context, id int) (models.Tour, error)
	UpdateTour(ctx context.Context, tour models.Tour) error
	DeleteTour(ctx context.Context, id int) error

	SetTitle(ctx context.Context, id int, title string) error
	SetStartTime(ctx context.Context, id int, startTime time.Time) error
	SetEndTime(ctx context.Context, id int, endTime time.Time) error
	SetLanguages(ctx context.Context, id int, languages string) error
	SetFreeCancellation(ctx context.Context, id int, freeCancellation bool) error
	SetCancellationCondition(ctx context.Context, id int, cancellationCondition string) error
	SetDescription(ctx context.Context, id int, description string) error
	SetMeetingPlace(ctx context.Context, id int, meetingPlace string) error
	SetArrivalPlace(ctx context.Context, id int, arrivalPlace string) error
	SetWhatIsIncluded(ctx context.Context, id int, whatIsIncluded string) error
	SetWhatToPrepare(ctx context.Context, id int, whatToPrepare string) error
	SetProhibitions(ctx context.Context, id int, prohibitions string) error
	SetPrice(ctx context.Context, id int, price int) error
	SetImages(ctx context.Context, id int, images string) error
}

type AuthTraveler interface {
	CreateTraveler(traveler models.TravelerSignUp) (int, error)
	GetTravelerByEmail(email string) (models.Traveler, error)
	GetAllTraveler() ([]models.Traveler, error)
	GetTravelerById(id int) (models.Traveler, error)
}

type Repository struct {
	Tour
	AuthTraveler
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Tour:         tour.NewTourPostgres(db),
		AuthTraveler: auth.NewAuthTavelerPostgres(db),
	}
}
