package repo

import (
	"context"
	"log"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"gorm.io/gorm"
	"tours.xws.com/model"
)

type TourRepository struct {
	DatabaseConnection *gorm.DB
}

// TOURS
func (repo *TourRepository) Get(id string, ctx context.Context) (model.Tour, error) {
	tracer := otel.Tracer("tours-repo")
	_, span := tracer.Start(ctx, "Get Repo")
	defer span.End()
	tour := model.Tour{}
	log.Println("Getting tour...")
	dbResult := repo.DatabaseConnection.First(&tour, "id = ?", id)
	if dbResult.Error != nil {
		return tour, dbResult.Error
	}
	return tour, nil
}

func (repo *TourRepository) GetAllByAuthor(authorId string, ctx context.Context) ([]model.Tour, error) {
	tracer := otel.Tracer("tours-repo")
	_, span := tracer.Start(ctx, "GetAuthors Repo")
	defer span.End()

	tours := []model.Tour{}
	log.Println("Getting all by author...")
	dbResult := repo.DatabaseConnection.Find(&tours, "author_id = ?", authorId)
	if dbResult.Error != nil {
		return tours, dbResult.Error
	}
	span.SetAttributes(attribute.Int("number_of_tours", len(tours)))
	return tours, nil
}

func (repo *TourRepository) CreateTour(tour *model.Tour) error {
	dbResult := repo.DatabaseConnection.Create(tour)
	log.Println("Creating tour...")
	if dbResult.Error != nil {
		return dbResult.Error
	}
	println("Rows affected: ", dbResult.RowsAffected)
	return nil
}

func (repo *TourRepository) UpdateTour(tour *model.Tour) error {
	dbResult := repo.DatabaseConnection.Save(tour)
	log.Println("Updating tour...")
	if dbResult.Error != nil {
		return dbResult.Error
	}
	return nil
}

// REQUIRED TIMES
func (repo *TourRepository) GetByTour(tourId string) (model.RequiredTime, error) {
	time := model.RequiredTime{}
	log.Println("Getting by tour...")
	dbResult := repo.DatabaseConnection.First(&time, "tour_id = ?", tourId)
	if dbResult != nil {
		return time, dbResult.Error
	}
	return time, nil
}

func (repo *TourRepository) GetAllByTour(tourId string) ([]model.RequiredTime, error) {
	time := []model.RequiredTime{}
	log.Println("Getting all by tour...")
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
	log.Println("Rows affected: ", dbResult.RowsAffected)
	return nil
}

func (repo *TourRepository) UpdateRequiredTime(time *model.RequiredTime) error {
	dbResult := repo.DatabaseConnection.Save(time)
	log.Println("Updating...")
	if dbResult.Error != nil {
		return dbResult.Error
	}
	return nil
}
