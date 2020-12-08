package controllers

import (
	"example/services"
	"example/structs"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(context *gin.Context) {
	var body structs.RegisterSchema
	if err = context.ShouldBindJSON(&body); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err = validate.Struct(body); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, err := services.Register(body)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"user": result})
}

func Login(context *gin.Context) {
	var body structs.LoginSchema
	if err = context.ShouldBindJSON(&body); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err = validate.Struct(body); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, err := services.Login(body)
	if err != nil {
		if err.Error() == "crypto/bcrypt: hashedPassword is not the hash of the given password" {
			context.JSON(http.StatusBadRequest, gin.H{"error": "Email or Password is wrong"})
		} else {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		return
	}
	context.JSON(http.StatusOK, gin.H{"user": result})
}
