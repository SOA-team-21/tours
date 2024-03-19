package service

import (
	"tours.xws.com/model"
	"tours.xws.com/repo"
)

type TourExecutionService struct {
	Repo         *repo.TourExecutionRepo
	KeyPointRepo *repo.KeyPointRepository
	TaskRepo     *repo.PointTaskRepo
}

func (service *TourExecutionService) Create(token *model.TourPurchaseToken) (*model.TourExecution, error) {
	execution := model.TourExecution{}
	execution.TourId = token.TourId
	execution.TouristId = token.TouristId
	if err := service.Repo.CreateExecution(&execution); err != nil {
		return nil, err
	}
	points, _ := service.KeyPointRepo.GetAllByTour(string(token.TourId))
	task := model.PointTask{}
	for _, point := range points {
		task.TourExecutionId = execution.Id
		task.KeyPointId = point.Id
		if err := service.TaskRepo.CreateTask(&task); err != nil {
			return nil, err
		}
	}
	return &execution, nil
}
