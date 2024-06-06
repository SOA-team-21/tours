package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"tours.xws.com/model"
	"tours.xws.com/proto/tours"
	"tours.xws.com/repo"
	"tours.xws.com/service"
)

const serviceName = "tours"

func initDB() *gorm.DB {
	dsn := "user=postgres password=super dbname=soa_tours host=tours-database port=5432 sslmode=disable"
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
	log.SetOutput(os.Stderr)

	// OpenTelemetry
	var err error
	tp, err = initTracer()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			log.Printf("Error shutting down tracer provider: %v", err)
		}
	}()
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))

	database := initDB()
	if database == nil {
		log.Println("FAILED TO CONNECT TO DB")
		return
	}

	keyPointRepo := &repo.KeyPointRepository{DatabaseConnection: database}
	tourRepo := &repo.TourRepository{DatabaseConnection: database}
	tourService := &service.TourService{Repo: tourRepo, KeyPointRepo: keyPointRepo}
	tourHandler := &TourHandler{TourService: tourService}

	lis, err := net.Listen("tcp", ":88")
	log.Println("Running gRPC on port 88")
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
	log.Println("Registered gRPC server")

	tours.RegisterToursServiceServer(grpcServer, tourHandler)
	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalln(err)
		}
	}()
	log.Println("Serving gRPC")

	stopCh := make(chan os.Signal)
	signal.Notify(stopCh, syscall.SIGTERM)

	<-stopCh

	grpcServer.Stop()
}
