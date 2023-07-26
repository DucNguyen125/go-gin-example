package services

import (
	"example/models"
	"example/structs"
	"example/utils/mysql"
	"time"
)

func CreateProduct(body structs.CreateProductSchema) (structs.Product, error) {
	newProduct := models.Product{
		ID:          body.ID,
		ProductCode: body.ProductCode,
		ProductName: body.ProductName,
		Price:       body.Price,
	}
	err := mysql.DB.Create(&newProduct)
	if err.Error != nil {
		return structs.Product{}, err.Error
	}
	createdProduct := structs.Product{
		ID:          newProduct.ID,
		ProductCode: newProduct.ProductCode,
		ProductName: newProduct.ProductName,
		Price:       newProduct.Price,
		CreatedAt:   newProduct.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   newProduct.UpdatedAt.Format(time.RFC3339),
	}
	return createdProduct, nil
}

func UpdateProduct(body structs.CreateProductSchema) (structs.Product, error) {
	product := structs.Product{}
	err := mysql.DB.Updates(models.Product{
		ID:          body.ID,
		ProductCode: body.ProductCode,
		ProductName: body.ProductName,
		Price:       body.Price,
	}).First(&product, body.ID).Error
	if err != nil {
		return structs.Product{}, err
	}
	return product, nil
}

func GetProduct(id string) (structs.Product, error) {
	product := structs.Product{}
	err := mysql.DB.Model(&models.Product{}).First(&product, id).Error
	if err != nil {
		return structs.Product{}, err
	}
	return product, nil
}

func GetListProduct(option structs.GetListProductSchema) []structs.Product {
	listProduct := []structs.Product{}
	limit := option.Limit
	skip := (option.Page - 1) * option.Limit
	mysql.DB.Model(&models.Product{}).Offset(skip).Limit(limit).Find(&listProduct)
	return listProduct
}

func DeleteProduct(id string) error {
	product := structs.Product{}
	err := mysql.DB.Model(&models.Product{}).Delete(&product, id).Error
	return err
}
