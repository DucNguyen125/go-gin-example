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
	ID          int
	ProductCode string
	ProductName string
	Price       int
}

type UpdateProductSchema struct {
	ProductCode string
	ProductName string
	Price       int
}

type GetListProductSchema struct {
	Page  int
	Limit int
}
