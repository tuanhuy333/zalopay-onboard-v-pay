package service

import (
	"errors"

	"gorm.io/gorm"

	"order-service/models"
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

func (s *Storage) GetAllOrder(p *[]models.Order) (*[]models.Order, error) {
	err := s.db.Find(&p)
	if err.Error != nil {
		return nil, err.Error
	}
	return p, nil
}

func (s *Storage) checkIfOrderExists(orderId int) bool {
	var o models.Order
	s.db.First(&o, orderId)
	if o.OrderNo == 0 {
		return false
	}
	return true
}

func (s *Storage) GetOrderById(orderId int) (*models.Order, error) {
	if s.checkIfOrderExists(orderId) == false {
		return nil, errors.New("No Order found")
	}
	var o models.Order
	s.db.First(&o, orderId)
	return &o, nil
}

func (s *Storage) UpdateOrderById(orderId int, o *models.Order) (*models.Order, error) {
	if s.checkIfOrderExists(orderId) == false {
		return nil, errors.New("No Order found")
	}
	o.Status = 1
	err := s.db.Save(&o)
	if err.Error != nil {
		return nil, err.Error
	}
	return nil, nil
}
