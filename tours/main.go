package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"tours.xws.com/handler"
	"tours.xws.com/model"
	"tours.xws.com/proto/tours"
	"tours.xws.com/repo"
	"tours.xws.com/service"
)

func initDB() *gorm.DB {
	dsn := "user=postgres password=super dbname=soa_tours host=localhost port=5432 sslmode=disable"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		print(err)
		return nil
	}
	err = database.AutoMigrate(&model.Tour{}, &model.KeyPoint{}, &model.TourExecution{}, &model.PointTask{}, &model.RequiredTime{}, &model.Preference{})
	if err != nil {
		log.Fatalf("Error migrating models: %v", err)
	}
	return database
}

func main() {
	database := initDB()
	if database == nil {
		print("FAILED TO CONNECT TO DB")
		return
	}

	keyPointRepo := &repo.KeyPointRepository{DatabaseConnection: database}
	tourRepo := &repo.TourRepository{DatabaseConnection: database}
	tourService := &service.TourService{Repo: tourRepo, KeyPointRepo: keyPointRepo}
	tourHandler := &handler.TourHandler{TourService: tourService}

	lis, err := net.Listen("tcp", ":88")
	fmt.Println("Running gRPC on port 88")
	if err != nil {
		log.Fatalln(err)
	}

	defer func(listener net.Listener) {
		err := listener.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(lis)

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	fmt.Println("Registered gRPC server")

	tours.RegisterToursServiceServer(grpcServer, tourHandler)
	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalln(err)
		}
	}()
	fmt.Println("Serving gRPC")

	stopCh := make(chan os.Signal)
	signal.Notify(stopCh, syscall.SIGTERM)

	<-stopCh

	grpcServer.Stop()
}
