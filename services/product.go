package services

import (
	"example/models"
	"example/utils/mysql"
	"time"
)

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

func CreateProduct(body CreateProductSchema) (Product, error) {
	newProduct := models.Product{
		ID:          body.ID,
		ProductCode: body.ProductCode,
		ProductName: body.ProductName,
		Price:       body.Price,
	}
	err := mysql.DB.Create(&newProduct)
	if err.Error != nil {
		return Product{}, err.Error
	}
	createdProduct := Product{
		ID:          newProduct.ID,
		ProductCode: newProduct.ProductCode,
		ProductName: newProduct.ProductName,
		Price:       newProduct.Price,
		CreatedAt:   newProduct.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   newProduct.UpdatedAt.Format(time.RFC3339),
	}
	return createdProduct, nil
}

func UpdateProduct(body CreateProductSchema) (Product, error) {
	product := Product{}
	err := mysql.DB.Updates(models.Product{
		ID:          body.ID,
		ProductCode: body.ProductCode,
		ProductName: body.ProductName,
		Price:       body.Price,
	}).First(&product, body.ID).Error
	if err != nil {
		return Product{}, err
	}
	return product, nil
}

func GetProduct(id string) (Product, error) {
	product := Product{}
	err := mysql.DB.Model(&models.Product{}).First(&product, id).Error
	if err != nil {
		return Product{}, err
	}
	return product, nil
}

func GetListProduct(option GetListProductSchema) []Product {
	listProduct := []Product{}
	limit := option.Limit
	skip := (option.Page - 1) * option.Limit
	mysql.DB.Model(&models.Product{}).Offset(skip).Limit(limit).Find(&listProduct)
	return listProduct
}

func DeleteProduct(id string) error {
	product := Product{}
	err := mysql.DB.Model(&models.Product{}).Delete(&product, id).Error
	return err
}
