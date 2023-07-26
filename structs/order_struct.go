package structs

type Order struct {
	ID          int    `json:"id"`
	OrderCode   string `json:"orderCode"`
	OrderType   string `json:"orderType"`
	Products    string `json:"products"`
	OrderStatus string `json:"orderStatus"`
	Quantity    int    `json:"quantity"`
	TotalPrice  int    `json:"totalPrice"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}

type CreateOrderSchema struct {
	ID          int
	OrderCode   string
	OrderType   string
	Products    string
	OrderStatus string
	Quantity    int
	TotalPrice  int
}

type UpdateOrderSchema struct {
	OrderCode   string
	OrderType   string
	Products    string
	OrderStatus string
	Quantity    int
	TotalPrice  int
}

type GetListOrderSchema struct {
	Page  int
	Limit int
}
