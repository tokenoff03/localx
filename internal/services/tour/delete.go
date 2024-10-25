package tour

import (
	"context"
)

func (t *TourService) DeleteTour(ctx context.Context, id int) error {
	return t.repo.DeleteTour(ctx, id)
}
