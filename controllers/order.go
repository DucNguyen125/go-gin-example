package controllers

import (
	"example/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateOrder(context *gin.Context) {
	var body services.CreateOrderSchema
	if err := context.ShouldBindJSON(&body); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := validate.Struct(body); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, err := services.CreateOrder(body)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"order": result})
}

func UpdateOrder(context *gin.Context) {
	var body services.UpdateOrderSchema
	if err := context.ShouldBindJSON(&body); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := validate.Struct(body); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, _ := strconv.Atoi(context.Param("id"))
	if err := validate.Var(id, "required,number"); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	dataUpdate := services.CreateOrderSchema{
		ID:          id,
		OrderCode:   body.OrderCode,
		OrderType:   body.OrderType,
		Products:    body.Products,
		OrderStatus: body.OrderStatus,
		Quantity:    body.Quantity,
		TotalPrice:  body.TotalPrice,
	}
	result, err := services.UpdateOrder(dataUpdate)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"order": result})
}

func GetOrder(context *gin.Context) {
	id := context.Param("id")
	if err := validate.Var(id, "required,number"); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, err := services.GetOrder(id)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"order": result})
}

func GetListOrder(context *gin.Context) {
	page, _ := strconv.Atoi(context.Query("page"))
	limit, _ := strconv.Atoi(context.Query("limit"))
	option := services.GetListOrderSchema{
		Page:  page,
		Limit: limit,
	}
	if err := validate.Struct(option); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result := services.GetListOrder(option)
	context.JSON(http.StatusOK, gin.H{"orders": result})
}

func DeleteOrder(context *gin.Context) {
	id := context.Param("id")
	if err := validate.Var(id, "required,number"); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := services.DeleteOrder(id); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"order": "Delete success"})
}
