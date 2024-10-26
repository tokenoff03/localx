package tour

import (
	"context"
)

func (t *TourService) DeleteTour(ctx context.Context, id int64) error {
	return t.repo.DeleteTour(ctx, id)
}
