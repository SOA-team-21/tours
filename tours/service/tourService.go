package service

import (
	"fmt"
	"time"

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
	populateTour(service, &Tour)
	return &Tour, nil
}

func (service *TourService) FindAllByAuthor(authorId string) ([]model.Tour, error) {
	tours, err := service.Repo.GetAllByAuthor(authorId)
	if err != nil {
		return nil, fmt.Errorf("cannot find tours by author with id %s", authorId)
	}
	for i := range tours {
		populateTour(service, &tours[i])
	}
	return tours, nil
}

func (service *TourService) Create(tour *model.Tour) (*model.Tour, error) {
	err := service.Repo.CreateTour(tour)
	if err != nil {
		return nil, err
	}
	return tour, nil
}

func (service *TourService) Update(tour *model.Tour) (*model.Tour, error) {
	err := service.Repo.UpdateTour(tour)
	if err != nil {
		return nil, err
	}
	// for i := range tour.KeyPoints {
	// 	service.KeyPointRepo.UpdateKeyPoint(&tour.KeyPoints[i])
	// }
	return tour, nil
}

func (service *TourService) Publish(tour *model.Tour) (*model.Tour, error) {
	fmt.Println("Status ture pre ažuriranja:", tour.Status)
	tour.Status = 1
	tour.PublishTime = time.Now()

	err := service.Repo.UpdateTour(tour)
	if err != nil {
		return nil, err
	}

	fmt.Println("Status ture posle ažuriranja:", tour.Status)
	return tour, nil
}

func (service *TourService) Archive(tour *model.Tour) (*model.Tour, error) {
	tour.Status = 2
	tour.PublishTime = time.Now()

	err := service.Repo.UpdateTour(tour)
	if err != nil {
		return nil, err
	}
	return tour, nil
}

func populateTour(service *TourService, Tour *model.Tour) { //Insertion of RequiredTimes and KeyPoints into Tour
	requiretTimesFromDb, err := service.Repo.GetAllByTour(fmt.Sprint(Tour.Id))
	if err == nil {
		Tour.RequiredTimes = append(Tour.RequiredTimes, requiretTimesFromDb...)
	}
	keyPointsFromDb, err := service.KeyPointRepo.GetAllByTour(fmt.Sprint(Tour.Id))
	if err == nil {
		Tour.KeyPoints = append(Tour.KeyPoints, keyPointsFromDb...)
	}
}
