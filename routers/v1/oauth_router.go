package v1

import (
	"example/controllers"

	"github.com/gin-gonic/gin"
)

func InitOAuthRouter(router *gin.RouterGroup) {
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
