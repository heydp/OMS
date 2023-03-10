package req

type Order struct {
	Id int ` json:"id" `
	// Base
	Status string    `json:"status"`
	Item   []AllItem `json:"items"`
}

type AllItem struct {
	Id int `json:"item_id"`

	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Quantity    int     `json:"quantity"`
}
