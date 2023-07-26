package routers

import (
	"example/middlewares"
	v1 "example/routers/v1"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.New()
	// router.Use(gin.Logger())
	router.Use(middlewares.Logger)
	router.Use(gin.Recovery())

	v1Router := router.Group("/v1")
	{
		orderPrefix := v1Router.Group("/order", middlewares.AuthenticationMiddleware)
		{
			v1.InitOrderRouter(orderPrefix)
		}
		productPrefix := v1Router.Group("/product", middlewares.AuthenticationMiddleware)
		{
			v1.InitProductRouter(productPrefix)
		}
		authPrefix := v1Router.Group("/auth")
		{
			v1.InitLocalAuthRouter(authPrefix)
			v1.InitOAuthRouter(authPrefix)
		}
	}

	return router
}
