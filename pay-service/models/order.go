package models

import "time"

type Order struct {
	OrderNo     int64
	MerchantID  string
	AppID       int32
	Status      int32
	Amount      int64
	ProductCode string
	Title       string
	Description string
	CreateTime  time.Time
}
