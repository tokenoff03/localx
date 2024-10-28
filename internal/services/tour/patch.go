package tour

import (
	"context"
	"localx/internal/models"
)

func (s *TourService) UpdateTourDetails(ctx context.Context, tour models.Tour) error {
	return s.repo.UpdateTourDetails(ctx, tour)
}
