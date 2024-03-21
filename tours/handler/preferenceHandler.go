package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"tours.xws.com/model"
	"tours.xws.com/service"
)

type PreferenceHandler struct {
	PreferenceService *service.PreferenceService
}

func (handler *PreferenceHandler) Create(writer http.ResponseWriter, req *http.Request) {
	var preference model.Preference
	err := json.NewDecoder(req.Body).Decode(&preference)

	if err != nil {
		println("Error while parsing json")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.PreferenceService.Create(&preference)
	if err != nil {
		println("Error while creating a new preference")
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(preference)
}

func (handler *PreferenceHandler) Update(writer http.ResponseWriter, req *http.Request) {
	var preference model.Preference
	err := json.NewDecoder(req.Body).Decode(&preference)

	if err != nil {
		println("Error while parsing json")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.PreferenceService.Update(&preference)
	if err != nil {
		println("Error while creating a new preference")
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(preference)
}

func (handler *PreferenceHandler) Delete(writer http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	preferenceID := params["preferenceId"]

	err := handler.PreferenceService.Delete(preferenceID)
	if err != nil {
		println(err.Error())
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(map[string]string{"message": "Preference deleted successfully"})
}

func (handler *PreferenceHandler) GetAllByUser(writer http.ResponseWriter, req *http.Request) {
	userId := mux.Vars(req)["userId"]
	println("UserId: %s", userId)
	preferences, err := handler.PreferenceService.GetAllByUser(userId)
	if err != nil {
		println(err.Error())
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(preferences)
}
