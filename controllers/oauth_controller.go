package controllers

import (
	"example/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/markbates/goth/gothic"
)

func LoginGoogle(context *gin.Context) {
	q := context.Request.URL.Query()
	q.Add("provider", "google")
	context.Request.URL.RawQuery = q.Encode()
	gothic.BeginAuthHandler(context.Writer, context.Request)
}

func LoginGoogleCallback(context *gin.Context) {
	result, err := services.LoginGoogleCallback(context)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"user": result})
}

func LoginFacebook(context *gin.Context) {
	q := context.Request.URL.Query()
	q.Add("provider", "facebook")
	context.Request.URL.RawQuery = q.Encode()
	gothic.BeginAuthHandler(context.Writer, context.Request)
}

func LoginFacebookCallback(context *gin.Context) {
	result, err := services.LoginFacebookCallback(context)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"user": result})
}
