package repo

import (
	"gorm.io/gorm"
	"tours.xws.com/model"
)

type TourRepository struct {
	DatabaseConnection *gorm.DB
}

// TOURS
func (repo *TourRepository) Get(id string) (model.Tour, error) {
	tour := model.Tour{}
	dbResult := repo.DatabaseConnection.First(&tour, "id = ?", id)
	if dbResult.Error != nil {
		return tour, dbResult.Error
	}
	return tour, nil
}

func (repo *TourRepository) GetAllByAuthor(authorId string) ([]model.Tour, error) {
	tours := []model.Tour{}
	dbResult := repo.DatabaseConnection.Find(&tours, "author_id = ?", authorId)
	if dbResult.Error != nil {
		return tours, dbResult.Error
	}
	return tours, nil
}

func (repo *TourRepository) CreateTour(tour *model.Tour) error {
	dbResult := repo.DatabaseConnection.Create(tour)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	println("Rows affected: ", dbResult.RowsAffected)
	return nil
}

func (repo *TourRepository) UpdateTour(tour *model.Tour) error {
	dbResult := repo.DatabaseConnection.Save(tour)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	return nil
}

// REQUIRED TIMES
func (repo *TourRepository) GetByTour(tourId string) (model.RequiredTime, error) {
	time := model.RequiredTime{}
	dbResult := repo.DatabaseConnection.First(&time, "tour_id = ?", tourId)
	if dbResult != nil {
		return time, dbResult.Error
	}
	return time, nil
}

func (repo *TourRepository) GetAllByTour(tourId string) ([]model.RequiredTime, error) {
	time := []model.RequiredTime{}
	dbResult := repo.DatabaseConnection.Find(&time, "tour_id = ?", tourId)
	if dbResult != nil {
		return time, dbResult.Error
	}
	return time, nil
}

func (repo *TourRepository) CreateRequiredTime(time *model.RequiredTime) error {
	dbResult := repo.DatabaseConnection.Create(time)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	println("Rows affected: ", dbResult.RowsAffected)
	return nil
}

func (repo *TourRepository) UpdateRequiredTime(time *model.RequiredTime) error {
	dbResult := repo.DatabaseConnection.Save(time)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	return nil
}
