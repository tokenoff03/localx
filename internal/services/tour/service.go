package tour

import (
	"localx/internal/repository"
)

type TourService struct {
	repo *repository.Repository
}

func NewTour(repos *repository.Repository) *TourService {
	return &TourService{repo: repos}
}
