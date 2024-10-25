package tour

import (
	"context"
	"localx/internal/models"
)

func (t *TourService) GetTourById(ctx context.Context, id int) (models.Tour, error) {
	return t.repo.GetTourById(ctx, id)
}
