package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"tours.xws.com/model"
	"tours.xws.com/service"
)

type TourExecutionHandler struct {
	Service *service.TourExecutionService
}

func (handler *TourExecutionHandler) Create(writer http.ResponseWriter, req *http.Request) {
	var token model.TourPurchaseToken
	err := json.NewDecoder(req.Body).Decode(&token)
	if err != nil {
		println("Error while parsing json")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	execution, err := handler.Service.Create(&token)
	if err != nil {
		println("Error while creating a new TourExecution")
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(execution)
}

func (handler *TourExecutionHandler) QuitExecution(writer http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	log.Printf("TourExecution sa id-em %s", id)
	execution, err := handler.Service.QuitExecution(id)
	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(execution)
}
