package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	Id int `gorm:"PRIMARY_KEY" json:"id" `
	// Base
	Status string    `json:"status"`
	Item   []AllItem `gorm:"foreignkey:OrderID;references:Id" json:"items"`
}

type AllItem struct {
	gorm.Model
	Id int `gorm:"PRIMARY_KEY" json:"item_id"`
	// OrderId     int     `gorm:"column:order_id" json: "order_id"`
	// Base
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Quantity    int     `json:"quantity"`
	OrderID     uint    `json:"order_id"`
}

// postgres.Jsonb `gorm:"type:jsonb;column:mdm_info" json:"mdm_info"`
