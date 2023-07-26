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
	ID          int    `validate:"required,number"`
	OrderCode   string `validate:"required,alphanum"`
	OrderType   string `validate:"required,alpha"`
	Products    string `validate:"required"`
	OrderStatus string `validate:"required,oneof='success' 'fail' 'pending'"`
	Quantity    int    `validate:"required,number"`
	TotalPrice  int    `validate:"required,number"`
}

type UpdateOrderSchema struct {
	OrderCode   string `validate:"alphanum"`
	OrderType   string `validate:"alpha"`
	Products    string
	OrderStatus string `validate:"oneof='success' 'fail' 'pending'"`
	Quantity    int    `validate:"number"`
	TotalPrice  int    `validate:"number"`
}

type GetListOrderSchema struct {
	Page  int `validate:"required,number,gte=1"`
	Limit int `validate:"number,gte=1,lte=100"`
}
