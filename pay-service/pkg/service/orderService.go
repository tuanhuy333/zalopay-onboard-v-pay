package service

import (
	"errors"

	"gorm.io/gorm"

	"V_Pay_Onboard_Program/models"
)

type Storage struct {
	db *gorm.DB
}

func New(db *gorm.DB) *Storage {
	return &Storage{db: db}
}
func (s *Storage) CreateOrder(p *models.Order) error {

	err := s.db.Create(&p)
	if err.Error != nil {
		return err.Error
	}
	return nil
}
func (s *Storage) GetAllOrder(p *[]models.Order) error {

	err := s.db.Find(&p)
	if err.Error != nil {
		return err.Error
	}
	return nil
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
