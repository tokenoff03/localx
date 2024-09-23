package service

import (
	"localx"
	"localx/pkg/repository"
)

type TourService struct {
	repo *repository.Repository
}

func NewTour(repos *repository.Repository) *TourService {
	return &TourService{repo: repos}
}

func (t *TourService) CreateTour(tour localx.Tour, companyId int) (int, error) {

	return 0, nil
}

func (t *TourService) GetById(id int) (localx.Tour, error) {

	return localx.Tour{}, nil
}
