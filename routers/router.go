package routers

import (
	v1 "example/routers/v1"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.New()

	v1Router := router.Group("/v1")
	{
		orderPrefix := v1Router.Group("/order")
		{
			v1.InitOrderRouter(orderPrefix)
		}
		productPrefix := v1Router.Group("/product")
		{
			v1.InitProductRouter(productPrefix)
		}
		localAuthPrefix := v1Router.Group("/auth")
		{
			v1.InitLocalAuthRouter(localAuthPrefix)
		}
	}

	return router
}
