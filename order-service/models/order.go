package models

import (
	"time"
)

type Order struct {
	OrderNo     int       `gorm:"primary_key;AUTO_INCREMENT" json:"orderNo"`
	MerchantID  string    `json:"merchantID"`
	AppID       int32     `json:"appID" binding:"required"`
	Status      int32     `json:"status"`
	Amount      float64   `json:"amount" binding:"required"`
	ProductCode string    `json:"productCode"`
	Description string    `json:"description"`
	CreateTime  time.Time `gorm:"autoCreateTime"`
	// this field in OrderReq but put in this for easy
	Mac string `gorm:"-" binding:"required"`
}
