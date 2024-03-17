package repo

import (
	"gorm.io/gorm"
	"tours.xws.com/model"
)

type PointTaskRepo struct {
	DatabaseConnection *gorm.DB
}

func (repo *PointTaskRepo) CreateTask(task *model.PointTask) error {
	dbResult := repo.DatabaseConnection.Create(task)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	return nil
}
