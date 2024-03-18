package handler

import (
	"encoding/json"
	"net/http"

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
