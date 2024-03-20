package service

import (
	"fmt"
	"time"

	"tours.xws.com/model"
	"tours.xws.com/repo"
	"tours.xws.com/utilities"
)

const _pointProximity int = 100

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
	for i, point := range points {
		task := model.PointTask{}
		task.Id = 0
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

func (service *TourExecutionService) QuitExecution(id string) (*model.TourExecution, error) {
	execution, err := service.Repo.GetExecution(id)
	if err != nil {
		return nil, err
	}
	execution.Status = model.Abandoned
	if err := service.Repo.Update(execution); err != nil {
		return nil, err
	}
	return execution, nil
}

func (service *TourExecutionService) UpdatePosition(currentPosition *model.Position, id string) (*model.TourExecution, error) {
	completedTasks := true
	execution, err := service.Repo.GetExecution(id)
	if err != nil {
		return nil, err
	}
	execution.Position = *currentPosition

	for i := range execution.Tasks {
		point, err := service.KeyPointRepo.Get(fmt.Sprint(execution.Tasks[i].KeyPointId))
		if err != nil { //TODO: Create different outcome for this
			break
		}
		if !execution.Tasks[i].Done {
			distance := utilities.CalculateDistance(point.Latitude, point.Longitude, *currentPosition)
			if distance*1000 <= float64(_pointProximity) {
				execution.Tasks[i].Done = true
				execution.Tasks[i].DoneOn = time.Now()
			} else {
				completedTasks = false //Finds one that is not done and not in proximity
			}
		}
	}
	if completedTasks {
		execution.Status = model.Completed
	}
	if err := service.Repo.Update(execution); err != nil {
		return nil, err
	}
	return execution, nil
}
