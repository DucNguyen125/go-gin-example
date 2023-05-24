package services

import (
	"example/models"
	"example/utils/mysql"
	"time"
)

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

func CreateOrder(body CreateOrderSchema) (Order, error) {
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
		return Order{}, err
	}
	createdOrder := Order{
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

func UpdateOrder(body CreateOrderSchema) (Order, error) {
	order := Order{}
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
		return Order{}, err
	}
	return order, nil
}

func GetOrder(id string) (Order, error) {
	order := Order{}
	err := mysql.DB.Model(&models.Order{}).First(&order, id).Error
	if err != nil {
		return Order{}, err
	}
	return order, nil
}

func GetListOrder(option GetListOrderSchema) []Order {
	listOrder := []Order{}
	limit := option.Limit
	skip := (option.Page - 1) * option.Limit
	mysql.DB.Model(&models.Order{}).Offset(skip).Limit(limit).Find(&listOrder)
	return listOrder
}

func DeleteOrder(id string) error {
	order := Order{}
	err := mysql.DB.Model(&models.Order{}).Delete(&order, id).Error
	return err
}
