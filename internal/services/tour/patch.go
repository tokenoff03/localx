package tour

import (
	"context"
	"time"
)

func (t *TourService) SetTitle(ctx context.Context, id int64, title string) error {
	return t.repo.SetTitle(ctx, id, title)
}

func (t *TourService) SetStartTime(ctx context.Context, id int64, startTime time.Time) error {
	return t.repo.SetStartTime(ctx, id, startTime)
}

func (t *TourService) SetEndTime(ctx context.Context, id int64, endTime time.Time) error {
	return t.repo.SetEndTime(ctx, id, endTime)
}

func (t *TourService) SetLanguages(ctx context.Context, id int64, languages string) error {
	return t.repo.SetLanguages(ctx, id, languages)
}

func (t *TourService) SetFreeCancellation(ctx context.Context, id int64, freeCancellation bool) error {
	return t.repo.SetFreeCancellation(ctx, id, freeCancellation)
}

func (t *TourService) SetCancellationCondition(ctx context.Context, id int64, cancellationCondition string) error {
	return t.repo.SetCancellationCondition(ctx, id, cancellationCondition)
}

func (t *TourService) SetDescription(ctx context.Context, id int64, description string) error {
	return t.repo.SetDescription(ctx, id, description)
}

func (t *TourService) SetMeetingPlace(ctx context.Context, id int64, meetingPlace string) error {
	return t.repo.SetMeetingPlace(ctx, id, meetingPlace)
}

func (t *TourService) SetArrivalPlace(ctx context.Context, id int64, arrivalPlace string) error {
	return t.repo.SetArrivalPlace(ctx, id, arrivalPlace)
}

func (t *TourService) SetWhatIsIncluded(ctx context.Context, id int64, whatIsIncluded string) error {
	return t.repo.SetWhatIsIncluded(ctx, id, whatIsIncluded)
}

func (t *TourService) SetWhatToPrepare(ctx context.Context, id int64, whatToPrepare string) error {
	return t.repo.SetWhatToPrepare(ctx, id, whatToPrepare)
}

func (t *TourService) SetProhibitions(ctx context.Context, id int64, prohibitions string) error {
	return t.repo.SetProhibitions(ctx, id, prohibitions)
}

func (t *TourService) SetPrice(ctx context.Context, id int64, price int64) error {
	return t.repo.SetPrice(ctx, id, price)
}

func (t *TourService) SetImages(ctx context.Context, id int64, images string) error {
	return t.repo.SetImages(ctx, id, images)
}
