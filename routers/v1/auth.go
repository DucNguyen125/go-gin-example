package v1

import (
	"example/controllers"

	"github.com/gin-gonic/gin"
)

func InitAuthRouter(router gin.IRouter) {
	router.POST("/register", func(context *gin.Context) {
		controllers.Register(context)
	})
	router.POST("/login", func(context *gin.Context) {
		controllers.Login(context)
	})
	router.GET("/google/login", func(context *gin.Context) {
		controllers.LoginGoogle(context)
	})
	router.GET("/google/callback", func(context *gin.Context) {
		controllers.LoginGoogleCallback(context)
	})
	router.GET("/facebook/login", func(context *gin.Context) {
		controllers.LoginFacebook(context)
	})
	router.GET("/facebook/callback", func(context *gin.Context) {
		controllers.LoginFacebookCallback(context)
	})
}
