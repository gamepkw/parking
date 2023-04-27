package repository

import "github.com/vidu171/clean-architecture-go/domain"

type DBHandler interface {
	GetAvailableParkings() ([]*domain.Parking, error)
	UpdateAvailableParking(book domain.Parking) error
}
