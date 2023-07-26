package routers

import (
	"example/middlewares"
	v1Routers "example/routers/v1"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.New()
	// router.Use(gin.Logger())
	router.Use(middlewares.Logger)
	router.Use(gin.Recovery())
	apiRouter := router.Group("/api")
	apiRouter.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	v1Routers.InitV1Router(apiRouter.Group("/v1", middlewares.Authentication))
	return router
}
