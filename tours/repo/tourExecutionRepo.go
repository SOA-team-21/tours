package repo

import (
	"errors"

	"gorm.io/gorm"
	"tours.xws.com/model"
)

type TourExecutionRepo struct {
	DatabaseConnection *gorm.DB
}

func (repo *TourExecutionRepo) CreateExecution(tourExecution *model.TourExecution) (*model.TourExecution, error) {
	existingExecution := model.TourExecution{}
	if err := repo.DatabaseConnection.Preload("Tasks").
		First(&existingExecution, "tour_id = ? AND tourist_id = ? AND status = ?", tourExecution.TourId, tourExecution.TouristId, 0).Error; err == nil {
		return &existingExecution, nil
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if err := repo.DatabaseConnection.Create(tourExecution).Error; err != nil {
		return nil, err
	}
	return tourExecution, nil
}

func (repo *TourExecutionRepo) GetExecution(id string) (*model.TourExecution, error) {
	execution := model.TourExecution{}
	dbResult := repo.DatabaseConnection.First(&execution, "id = ?", id)
	if dbResult.Error != nil {
		return nil, dbResult.Error
	}
	return &execution, nil
}

func (repo *TourExecutionRepo) Update(execution *model.TourExecution) error {
	dbResult := repo.DatabaseConnection.Save(execution)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	return nil
}
