package v1

import (
	"example/controllers"

	"github.com/gin-gonic/gin"
)

func InitLocalAuthRouter(router *gin.RouterGroup) {
	router.POST("/register", func(context *gin.Context) {
		controllers.Register(context)
	})
	router.POST("/login", func(context *gin.Context) {
		controllers.Login(context)
	})
}
