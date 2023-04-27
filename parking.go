package domain

type Parking struct {
	Parkingid       int64  `json:"parkingid"`
	Parkingtype     string `json:"parkingtype"`
	Is_occupied     string `json:"is_occupied" validate:"required"`
	Lastestcheckin  string `json:"lastestcheckin" validate:"required"`
	Lastestcheckout string `json:"lastestcheckout"`
}

type Available struct {
	ID    int64 `json:"id"`
	Floor int64 `json:"floor"`
	// CarparkingAvailable_1   int   `json:"carparkingAvailable_1"`
	// CarparkingAvailable_2   int   `json:"carparkingAvailable_2"`
	// CarparkingAvailable_3   int   `json:"carparkingAvailable_3"`
	// CarparkingAvailable_4   int   `json:"carparkingAvailable_4"`
	// CarparkingAvailable_5   int   `json:"carparkingAvailable_5"`
	// MotorparkingAvailable_1 int   `json:"motorparkingAvailable_1"`
	// MotorparkingAvailable_2 int   `json:"motorparkingAvailable_2"`
}
