package service

import (
	"fmt"

	"tours.xws.com/model"
	"tours.xws.com/repo"
)

type TourService struct {
	Repo         *repo.TourRepository
	KeyPointRepo *repo.KeyPointRepository
}

func (service *TourService) FindTour(id string) (*model.Tour, error) {
	Tour, err := service.Repo.Get(id)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("menu item with id %s not found", id))
	}

	//TODO: Insert RequiredTimes

	//Insertion of the KeyPoints, manually
	keyPointsFromDb, err := service.KeyPointRepo.GetAllByTour(Tour.Id.String())
	if err == nil {
		Tour.KeyPoints = append(Tour.KeyPoints, keyPointsFromDb...)
	}
	return &Tour, nil
}

func (service *TourService) Create(tour *model.Tour) error {
	err := service.Repo.CreateTour(tour)
	if err != nil {
		return err
	}
	return nil
}

func (service *TourService) Update(tour *model.Tour) error {
	err := service.Repo.UpdateTour(tour)
	if err != nil {
		return err
	}
	for i := range tour.KeyPoints {
		service.KeyPointRepo.UpdateKeyPoint(&tour.KeyPoints[i])
	}
	return nil
}
