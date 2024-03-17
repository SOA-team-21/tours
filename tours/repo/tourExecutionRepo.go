package repo

import (
	"gorm.io/gorm"
	"tours.xws.com/model"
)

type TourExecutionRepo struct {
	DatabaseConnection *gorm.DB
}

func (repo *TourExecutionRepo) CreateExecution(tourExecution *model.TourExecution) error {
	dbResult := repo.DatabaseConnection.Create(tourExecution)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	println("Rows affected: ", dbResult.RowsAffected)
	return nil
}
