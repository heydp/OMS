package dto

import "github.com/heydp/oms/pkg/models"

type Order struct {
	Id int ` json:"id" `
	// Base
	Status   string    `json:"status"`
	Total    float64   `json:"total"`
	Currency string    `json:"currencyUnit"`
	Item     []AllItem `json:"items"`
}

type AllItem struct {
	Id int `json:"item_id"`

	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Quantity    int     `json:"quantity"`
}

func Convert(op *Order, mod *models.Order) {
	op.Id = mod.Id
	op.Status = mod.Status
	op.Currency = "USD"
	var total float64
	total = 0
	for _, val := range mod.Item {
		var allItem AllItem
		allItem.Id = val.Id
		allItem.Description = val.Description
		allItem.Price = val.Price
		allItem.Quantity = val.Quantity
		total = total + allItem.Price*float64(allItem.Quantity)
		op.Item = append(op.Item, allItem)
	}
	op.Total = total
}
