package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"tours.xws.com/model"
	"tours.xws.com/service"
)

type KeyPointHandler struct {
	KeyPointService *service.KeyPointService
}

func (handler *KeyPointHandler) Get(writer http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	log.Printf("KeyPoint sa id-em %s", id)
	point, err := handler.KeyPointService.FindKeyPoint(id)
	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(point)
}

func (handler *KeyPointHandler) GetAllByTour(writer http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["tourId"]
	log.Printf("Tour sa id-em %s", id)
	points, err := handler.KeyPointService.FindAllByTour(id)
	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(points)
}

func (handler *KeyPointHandler) Create(writer http.ResponseWriter, req *http.Request) {
	var keyPoint model.KeyPoint
	err := json.NewDecoder(req.Body).Decode(&keyPoint)
	if err != nil {
		println("Error while parsing json")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.KeyPointService.Create(&keyPoint)
	if err != nil {
		println("Error while creating a new KeyPoint")
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
}
