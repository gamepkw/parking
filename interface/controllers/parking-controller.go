package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/vidu171/clean-architecture-go/domain"

	"github.com/vidu171/clean-architecture-go/usecases"
)

type ParkingController struct {
	parkingInteractor usecases.ParkingInteractor
}

func NewParkingController(parkingInteractor usecases.ParkingInteractor) *ParkingController {
	return &ParkingController{parkingInteractor}
}

func (controller *ParkingController) Add(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	var parking domain.Parking
	err := json.NewDecoder(req.Body).Decode(&parking)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(ErrorResponse{Message: "Invalid Payload"})
		return
	}
	err2 := controller.parkingInteractor.CreateParking(parking)
	if err2 != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(ErrorResponse{Message: err2.Error()})
		return
	}
	res.WriteHeader(http.StatusOK)
}

func (controller *ParkingController) FindAll(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	results, err2 := controller.parkingInteractor.FindAll()
	if err2 != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(ErrorResponse{Message: err2.Error()})
		return
	}
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(results)
}
