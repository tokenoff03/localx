package tour

import (
	"context"
	"localx/internal/models"
)

func (t *TourService) CreateTour(ctx context.Context, tour models.Tour, companyId int) (int, error) {
	return t.repo.CreateTour(ctx, tour, companyId)
}
