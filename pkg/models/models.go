package models

type Order struct {
	// gorm.Model
	// Base `json:"-"`
	Id int `gorm:"PRIMARY_KEY uniqueIndex" json:"id" `

	Status string    `json:"status"`
	Item   []AllItem `gorm:"foreignkey:OrderID;references:Id" json:"items"`
}

type AllItem struct {
	// gorm.Model
	// ItemBase `json:"-"`
	Id int ` gorm:"PRIMARY_KEY" sql:"AUTO_INCREMENT" json:"item_id"`

	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Quantity    int     `json:"quantity"`
	OrderID     uint    `json:"order_id"`
}

// postgres.Jsonb `gorm:"type:jsonb;column:mdm_info" json:"mdm_info"`
