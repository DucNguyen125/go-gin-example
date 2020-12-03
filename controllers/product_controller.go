package controllers

import (
	"example/services"
	"example/structs"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateProduct(context *gin.Context) {
	var body structs.CreateProductSchema
	err := context.ShouldBindJSON(&body)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var result structs.Product
	result, err = services.CreateProduct(body)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"product": result})
}

func UpdateProduct(context *gin.Context) {
	var body structs.UpdateProductSchema
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
	dataUpdate := structs.CreateProductSchema{
		ID:          id,
		ProductCode: body.ProductCode,
		ProductName: body.ProductName,
		Price:       body.Price,
	}
	var result structs.Product
	result, err = services.UpdateProduct(dataUpdate)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"product": result})
}

func GetProduct(context *gin.Context) {
	id := context.Param("id")
	result, err := services.GetProduct(id)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"product": result})
}

func GetListProduct(context *gin.Context) {
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
	option := structs.GetListProductSchema{
		Page:  page,
		Limit: limit,
	}
	result := services.GetListProduct(option)
	context.JSON(http.StatusOK, gin.H{"products": result})
}

func DeleteProduct(context *gin.Context) {
	id := context.Param("id")
	err := services.DeleteProduct(id)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"product": "Delete success"})
}
