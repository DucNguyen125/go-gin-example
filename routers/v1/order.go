package v1

import (
	"example/controllers"

	"github.com/gin-gonic/gin"
)

func InitOrderRouter(router gin.IRouter) {
	router.POST("/", func(context *gin.Context) {
		controllers.CreateOrder(context)
	})
	router.PUT("/:id", func(context *gin.Context) {
		controllers.UpdateOrder(context)
	})
	router.GET("/:id", func(context *gin.Context) {
		controllers.GetOrder(context)
	})
	router.GET("", func(context *gin.Context) {
		controllers.GetListOrder(context)
	})
	router.DELETE("/:id", func(context *gin.Context) {
		controllers.DeleteOrder(context)
	})
}
