package usecases

import (
	"main/domain"
	"math"
)

type BookInteractor struct {
	BookRepository domain.BookRepository
}

func feeCal(parkingTime vehicleType) (fee int) {
	if parkingTime <= 30 {
		fee = 0
	} else if parkingTime > 60 && parkingTime <= 240 {
		if Vehicletype == "Car" {
			fee = math.Ceil(parkingTime/60) * carparkingfee_14
		}
		if Vehicletype == "Motor" {
			fee = math.Ceil(parkingTime/60) * motorparkingfee_14
		}
	} else if parkingTime > 240 && parkingTime <= 720 {
		if Vehicletype == "Car" {
			fee = math.Ceil(parkingTime/60) * carparkingfee_412
		}
		if Vehicletype == "Motor" {
			fee = math.Ceil(parkingTime/60) * motorparkingfee_412
		}
		fee = math.Ceil(parkingTime/60) * carparkingfee_412
	} else if parkingTime > 720 && parkingTime <= 1440 {
		if Vehicletype == "Car" {
			fee = carparkingfee_allday
		}
		if Vehicletype == "Motor" {
			fee = motorparkingfee_allday
		}
	}
	return fee
}

func entryPoint(Floor Vehicletype) {
	if Vehicletype == "Car" {
		switch Floor {
		case 1:
			carparkingAvailable_1 - 1
		case 2:
			carparkingAvailable_2 - 1
		case 3:
			carparkingAvailable_3 - 1
		case 4:
			carparkingAvailable_4 - 1
		case 5:
			carparkingAvailable_5 - 1
		}
	} else if Vehicletype == "Motor" {
		switch Floor {
		case 1:
			motorparkingAvailable_1 - 1
		case 2:
			motorparkingAvailable_2 - 1
		}
	}
}

func exitPoint(Floor Vehicletype) {
	if vehicleType == "Car" {
		switch Floor {
		case 1:
			carparkingAvailable_1 + 1
		case 2:
			carparkingAvailable_2 + 1
		case 3:
			carparkingAvailable_3 + 1
		case 4:
			carparkingAvailable_4 + 1
		case 5:
			carparkingAvailable_5 + 1
		}
	} else if vehicleType == "Motor" {
		switch floor {
		case 1:
			motorparkingAvailable_1 + 1
		case 2:
			motorparkingAvailable_2 + 1
		}
	}
}
