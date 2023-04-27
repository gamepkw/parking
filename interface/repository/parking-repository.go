package repository

import (
	"github.com/vidu171/clean-architecture-go/domain"
)

type ParkingRepo struct {
	handler DBHandler
}

func NewBookRepo(handler DBHandler) ParkingRepo {
	return ParkingRepo{handler}
}

func (repo ParkingRepo) UpdateAvailableParking(book domain.Book) error {
	err := repo.handler.UpdateAvailableParking(book)
	if err != nil {
		return err
	}
	return nil
}

func (repo ParkingRepo) FindAll() ([]*domain.Book, error) {
	results, err := repo.handler.GetAvailableParkings()
	if err != nil {
		return results, err
	}
	return results, nil
}
