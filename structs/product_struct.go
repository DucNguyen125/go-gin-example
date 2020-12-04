package structs

type Product struct {
	ID          int    `json:"id"`
	ProductCode string `json:"productCode"`
	ProductName string `json:"productName"`
	Price       int    `json:"price"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}

type CreateProductSchema struct {
	ID          int    `validate:"required,number"`
	ProductCode string `validate:"required,alphanum"`
	ProductName string `validate:"required,alphanum"`
	Price       int    `validate:"required,number"`
}

type UpdateProductSchema struct {
	ProductCode string `validate:"alphanum"`
	ProductName string `validate:"alphanum"`
	Price       int    `validate:"number"`
}

type GetListProductSchema struct {
	Page  int `validate:"required,number,gte=1"`
	Limit int `validate:"required,number,gte=1,lte=100"`
}
