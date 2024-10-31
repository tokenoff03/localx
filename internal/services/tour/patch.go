package tour

import (
	"context"
	"localx/internal/models"
)

func (s *TourService) PartialUpdateTour(ctx context.Context, tour models.Tour) error {
	return s.repo.PartialUpdateTour(ctx, tour)
}
