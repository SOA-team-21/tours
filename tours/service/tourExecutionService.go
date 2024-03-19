package service

import (
	"fmt"

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

	executionFromDb, err := service.Repo.CreateExecution(&execution)
	if executionFromDb != nil && len(executionFromDb.Tasks) > 0 { //Execution already exsist
		return executionFromDb, nil
	}
	if err != nil { //Something wrong with creation of Execution
		return nil, err
	}

	// Velika nuzda -> sve ok
	points, _ := service.KeyPointRepo.GetAllByTour(fmt.Sprint(token.TourId))
	tasks := make([]model.PointTask, len(points))
	task := model.PointTask{}
	for i, point := range points {
		task.TourExecutionId = execution.Id
		task.KeyPointId = point.Id
		if err := service.TaskRepo.CreateTask(&task); err != nil {
			return nil, err
		}
		tasks[i] = task
	}
	execution.Tasks = tasks
	return &execution, nil
}
