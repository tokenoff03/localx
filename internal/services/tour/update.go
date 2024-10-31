package tour

import (
	"context"
	"localx/internal/models"
)

func (t *TourService) UpdateTour(ctx context.Context, tour models.Tour) error {
	return t.repo.UpdateTour(ctx, tour)
}
