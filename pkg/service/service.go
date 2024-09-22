package service

import (
	"localx"
	"localx/pkg/repository"
)

type Tour interface {
	CreateTour(tour localx.Tour, companyId int) (int, error)
	GetById(id int) (localx.Tour, error)
}

type Service struct {
	Tour
}

func NewService(repo *repository.Repository) *Service {
	return &Service{}
}
