package models

import (
	"time"
)

type Order struct {
	OrderNo     int       `gorm:"primary_key;AUTO_INCREMENT" json:"orderNo"`
	MerchantID  string    `json:"merchantID"`
	AppID       int32     `json:"appID"`
	Status      int32     `json:"status"`
	Amount      float64   `json:"amount"`
	ProductCode string    `json:"productCode"`
	Description string    `json:"description"`
	CreateTime  time.Time `gorm:"autoCreateTime"`
}
