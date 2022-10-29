package service

import (
	"errors"

	"gorm.io/gorm"

	"pay-service/models"
)

type Storage struct {
	db *gorm.DB
}

func New(db *gorm.DB) *Storage {
	return &Storage{db: db}
}
func (s *Storage) checkIfOrderExists(orderId string) bool {
	var o models.Order
	s.db.First(&o, orderId)
	if o.OrderNo == 0 {
		return false
	}
	return true
}

func (s *Storage) GetOrderById(orderId string) (*models.Order, error) {
	if s.checkIfOrderExists(orderId) == false {
		return nil, errors.New("No Order found")
	}
	var o models.Order
	s.db.First(&o, orderId)
	return &o, nil
}
