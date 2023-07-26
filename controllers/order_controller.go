package controllers

import (
	"example/services"
	"example/structs"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateOrder(context *gin.Context) {
	var body structs.CreateOrderSchema
	err := context.ShouldBindJSON(&body)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var result structs.Order
	result, err = services.CreateOrder(body)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"order": result})
}

func UpdateOrder(context *gin.Context) {
	var body structs.UpdateOrderSchema
	var err error
	err = context.ShouldBindJSON(&body)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var id int
	id, err = strconv.Atoi(context.Param("id"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	dataUpdate := structs.CreateOrderSchema{
		ID:          id,
		OrderCode:   body.OrderCode,
		OrderType:   body.OrderType,
		Products:    body.Products,
		OrderStatus: body.OrderStatus,
		Quantity:    body.Quantity,
		TotalPrice:  body.TotalPrice,
	}
	var result structs.Order
	result, err = services.UpdateOrder(dataUpdate)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"order": result})
}

func GetOrder(context *gin.Context) {
	id := context.Param("id")
	result, err := services.GetOrder(id)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"order": result})
}

func GetListOrder(context *gin.Context) {
	var err error
	var page int
	var limit int
	page, err = strconv.Atoi(context.Query("page"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	limit, err = strconv.Atoi(context.Query("limit"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	option := structs.GetListOrderSchema{
		Page:  page,
		Limit: limit,
	}
	result := services.GetListOrder(option)
	context.JSON(http.StatusOK, gin.H{"orders": result})
}

func DeleteOrder(context *gin.Context) {
	id := context.Param("id")
	err := services.DeleteOrder(id)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"order": "Delete success"})
}
