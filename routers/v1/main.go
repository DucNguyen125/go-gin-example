package v1

import (
	"github.com/gin-gonic/gin"
)

func InitV1Router(
	r *gin.RouterGroup,
) {
	r.Use()
	{
		InitAuthRouter(r.Group("/auth"))
		InitProductRouter(r.Group("/products"))
		InitOrderRouter(r.Group("/orders"))
	}
}
