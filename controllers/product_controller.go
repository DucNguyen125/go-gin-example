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
	if err = context.ShouldBindJSON(&body); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err = validate.Struct(body); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, err := services.CreateProduct(body)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"product": result})
}

func UpdateProduct(context *gin.Context) {
	var body structs.UpdateProductSchema
	if err = context.ShouldBindJSON(&body); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err = validate.Struct(body); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, _ := strconv.Atoi(context.Param("id"))
	if err = validate.Var(id, "required,number"); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	dataUpdate := structs.CreateProductSchema{
		ID:          id,
		ProductCode: body.ProductCode,
		ProductName: body.ProductName,
		Price:       body.Price,
	}
	result, err := services.UpdateProduct(dataUpdate)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"product": result})
}

func GetProduct(context *gin.Context) {
	id := context.Param("id")
	if err = validate.Var(id, "required,number"); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, err := services.GetProduct(id)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"product": result})
}

func GetListProduct(context *gin.Context) {
	page, _ := strconv.Atoi(context.Query("page"))
	limit, _ := strconv.Atoi(context.Query("limit"))
	option := structs.GetListProductSchema{
		Page:  page,
		Limit: limit,
	}
	if err = validate.Struct(option); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result := services.GetListProduct(option)
	context.JSON(http.StatusOK, gin.H{"products": result})
}

func DeleteProduct(context *gin.Context) {
	id := context.Param("id")
	if err = validate.Var(id, "required,number"); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err = services.DeleteOrder(id); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"product": "Delete success"})
}
