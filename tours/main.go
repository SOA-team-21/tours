package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"tours.xws.com/handler"
	"tours.xws.com/model"
	"tours.xws.com/repo"
	"tours.xws.com/service"
)

func initDB() *gorm.DB {
	dsn := "user=postgres password=super dbname=soa_tours host=localhost port=5432 sslmode=disable search_path=tours"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		print(err)
		return nil
	}
	err = database.AutoMigrate(&model.Tour{}, &model.KeyPoint{}, &model.TourExecution{}, &model.PointTask{}, &model.RequiredTime{})
	if err != nil {
		log.Fatalf("Error migrating models: %v", err)
	}
	return database
}

func startServer(handler *handler.TourHandler, keyPointHandler *handler.KeyPointHandler, tourExeHandler *handler.TourExecutionHandler) {
	router := mux.NewRouter().StrictSlash(true)

	//TOURS
	router.HandleFunc("/tour/create", handler.Create).Methods("POST")
	router.HandleFunc("/tour/getTour/{id}", handler.Get).Methods("GET")
	router.HandleFunc("/tour/getAllByAuthor/{id}", handler.GetAllByAuthor).Methods("GET")
	router.HandleFunc("/tour/update", handler.Update).Methods("PUT")

	//KEYPOINTS
	router.HandleFunc("/keypoint/create", keyPointHandler.Create).Methods("POST")
	router.HandleFunc("/keypoint/getKeyPoint/{id}", keyPointHandler.Get).Methods("GET")
	router.HandleFunc("/keypoint/getAllByTour/{tourId}", keyPointHandler.GetAllByTour).Methods("GET")

	//TOUREXECUTIONS
	router.HandleFunc("/tourexecution/create", tourExeHandler.Create).Methods("POST")

	println("Server starting")
	log.Fatal(http.ListenAndServe(":88", router)) //Port number must be different for different servers (because all run on localhost)
}

func main() {
	database := initDB()
	if database == nil {
		print("FAILED TO CONNECT TO DB")
		return
	}

	keyPointRepo := &repo.KeyPointRepository{DatabaseConnection: database}
	keyPointService := &service.KeyPointService{Repo: keyPointRepo}
	keyPointHandler := &handler.KeyPointHandler{KeyPointService: keyPointService}

	tourRepo := &repo.TourRepository{DatabaseConnection: database}
	tourService := &service.TourService{Repo: tourRepo, KeyPointRepo: keyPointRepo}
	tourHandler := &handler.TourHandler{TourService: tourService}

	pointTaskRepo := &repo.PointTaskRepo{DatabaseConnection: database}

	tourExecutionRepo := &repo.TourExecutionRepo{DatabaseConnection: database}
	tourExecutionService := &service.TourExecutionService{Repo: tourExecutionRepo, KeyPointRepo: keyPointRepo, TaskRepo: pointTaskRepo}
	tourExecutionHandler := &handler.TourExecutionHandler{Service: tourExecutionService}

	startServer(tourHandler, keyPointHandler, tourExecutionHandler)
}
