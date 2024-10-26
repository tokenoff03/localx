package services

import (
	"context"
	"localx/internal/models"
	"localx/internal/repository"
	"localx/internal/services/auth"
	"localx/internal/services/tour"
	"time"
)

type Tour interface {
	CreateTour(ctx context.Context, tour models.Tour, companyId int64) (int64, error)
	GetTourById(ctx context.Context, id int64) (models.Tour, error)
	UpdateTour(ctx context.Context, tour models.Tour) error
	DeleteTour(ctx context.Context, id int64) error

	SetTitle(ctx context.Context, id int64, title string) error
	SetStartTime(ctx context.Context, id int64, startTime time.Time) error
	SetEndTime(ctx context.Context, id int64, endTime time.Time) error
	SetLanguages(ctx context.Context, id int64, languages string) error
	SetFreeCancellation(ctx context.Context, id int64, freeCancellation bool) error
	SetCancellationCondition(ctx context.Context, id int64, cancellationCondition string) error
	SetDescription(ctx context.Context, id int64, description string) error
	SetMeetingPlace(ctx context.Context, id int64, meetingPlace string) error
	SetArrivalPlace(ctx context.Context, id int64, arrivalPlace string) error
	SetWhatIsIncluded(ctx context.Context, id int64, whatIsIncluded string) error
	SetWhatToPrepare(ctx context.Context, id int64, whatToPrepare string) error
	SetProhibitions(ctx context.Context, id int64, prohibitions string) error
	SetPrice(ctx context.Context, id int64, price int64) error
	SetImages(ctx context.Context, id int64, images string) error
}

type AuthTraveler interface {
	CreateTraveler(traveler models.TravelerSignUp) (int, error)
	GetTravelerById(id int) (models.Traveler, error)
	GetTravelerByEmail(email string) (models.Traveler, error)
	GenerateToken(id int) (string, error)
	ParseToken(accessToken string) (string, error)
	GetAllTraveler() ([]models.Traveler, error)
	GenerateRefreshToken(id int) (string, error)
	GenerateAndSaveOTP(email string) (string, error)
	ValidateOTP(email, inputCode string) (bool, error)
	CleanExpiredOTPs()
	SendEmail(to string, subject string, body string) error
	GetTokens(phoneNumber string) (string, string, error)
	StoreTokens(phoneNumber, accessToken, refreshToken string)
	UpdateTokens(phoneNumber, accessToken, refreshToken string) error
}

type Services struct {
	Tour
	AuthTraveler
}

func NewServices(repo *repository.Repository) *Services {
	return &Services{
		Tour:         tour.NewTour(repo),
		AuthTraveler: auth.NewAuthTravelerService(repo),
	}
}
