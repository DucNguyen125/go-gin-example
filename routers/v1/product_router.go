package v1

import (
	"example/controllers"

	"github.com/gin-gonic/gin"
)

func InitProductRouter(router *gin.RouterGroup) {
	router.POST("/", func(context *gin.Context) {
		controllers.CreateProduct(context)
	})
	router.PUT("/:id", func(context *gin.Context) {
		controllers.UpdateProduct(context)
	})
	router.GET("/:id", func(context *gin.Context) {
		controllers.GetProduct(context)
	})
	router.GET("", func(context *gin.Context) {
		controllers.GetListProduct(context)
	})
	router.DELETE("/:id", func(context *gin.Context) {
		controllers.DeleteProduct(context)
	})
}
