package service

import (
	"fmt"

	"tours.xws.com/model"
	"tours.xws.com/repo"
)

type KeyPointService struct {
	Repo *repo.KeyPointRepository
}

func (service *KeyPointService) FindKeyPoint(id string) (*model.KeyPoint, error) {
	KeyPoint, err := service.Repo.Get(id)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("menu item with id %s not found", id))
	}
	return &KeyPoint, nil
}

func (service *KeyPointService) FindAllByTour(tourId string) ([]model.KeyPoint, error) {
	points, err := service.Repo.GetAllByTour(tourId)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("menu item with id %s not found", tourId))
	}
	return points, nil
}

func (service *KeyPointService) Create(keyPoint *model.KeyPoint) error {
	err := service.Repo.CreateKeyPoint(keyPoint)
	if err != nil {
		return err
	}
	return nil
}
