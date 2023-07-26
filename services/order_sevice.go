package services

import (
	"example/models"
	"example/structs"
	"example/utils/mysql"
	"time"
)

func CreateOrder(body structs.CreateOrderSchema) (structs.Order, error) {
	newOrder := models.Order{
		ID:          body.ID,
		OrderCode:   body.OrderCode,
		OrderType:   body.OrderType,
		Products:    body.Products,
		OrderStatus: body.OrderStatus,
		Quantity:    body.Quantity,
		TotalPrice:  body.TotalPrice,
	}
	if err := mysql.DB.Create(&newOrder).Error; err != nil {
		return structs.Order{}, err
	}
	createdOrder := structs.Order{
		ID:          newOrder.ID,
		OrderCode:   newOrder.OrderCode,
		OrderType:   newOrder.OrderType,
		Products:    newOrder.Products,
		OrderStatus: newOrder.OrderStatus,
		Quantity:    newOrder.Quantity,
		TotalPrice:  newOrder.TotalPrice,
		CreatedAt:   newOrder.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   newOrder.UpdatedAt.Format(time.RFC3339),
	}
	return createdOrder, nil
}

func UpdateOrder(body structs.CreateOrderSchema) (structs.Order, error) {
	order := structs.Order{}
	err := mysql.DB.Updates(models.Order{
		ID:          body.ID,
		OrderCode:   body.OrderCode,
		OrderType:   body.OrderType,
		Products:    body.Products,
		OrderStatus: body.OrderStatus,
		Quantity:    body.Quantity,
		TotalPrice:  body.TotalPrice,
	}).First(&order, body.ID).Error
	if err != nil {
		return structs.Order{}, err
	}
	return order, nil
}

func GetOrder(id string) (structs.Order, error) {
	order := structs.Order{}
	err := mysql.DB.Model(&models.Order{}).First(&order, id).Error
	if err != nil {
		return structs.Order{}, err
	}
	return order, nil
}

func GetListOrder(option structs.GetListOrderSchema) []structs.Order {
	listOrder := []structs.Order{}
	limit := option.Limit
	skip := (option.Page - 1) * option.Limit
	mysql.DB.Model(&models.Order{}).Offset(skip).Limit(limit).Find(&listOrder)
	return listOrder
}

func DeleteOrder(id string) error {
	order := structs.Order{}
	err := mysql.DB.Model(&models.Order{}).Delete(&order, id).Error
	return err
}
