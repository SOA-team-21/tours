package repo

import (
	"gorm.io/gorm"
	"tours.xws.com/model"
)

type KeyPointRepository struct {
	DatabaseConnection *gorm.DB
}

func (repo *KeyPointRepository) Get(id string) (model.KeyPoint, error) {
	point := model.KeyPoint{}
	dbResult := repo.DatabaseConnection.First(&point, "id = ?", id)
	if dbResult != nil {
		return point, dbResult.Error
	}
	return point, nil
}

func (repo *KeyPointRepository) GetAllByTour(tourId string) ([]model.KeyPoint, error) {
	points := []model.KeyPoint{}
	dbResult := repo.DatabaseConnection.Find(&points, "tour_id = ?", tourId)
	if dbResult != nil {
		return points, dbResult.Error
	}
	return points, nil
}

func (repo *KeyPointRepository) CreateKeyPoint(point *model.KeyPoint) error {
	dbResult := repo.DatabaseConnection.Create(point)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	println("Rows affected: ", dbResult.RowsAffected)
	return nil
}

func (repo *KeyPointRepository) UpdateKeyPoint(point *model.KeyPoint) error {
	dbResult := repo.DatabaseConnection.Save(point)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	return nil
}
