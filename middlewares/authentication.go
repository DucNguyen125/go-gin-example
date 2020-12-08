package middlewares

import (
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()
var err error

type Header struct {
	Authorization string `header:"authorization" validate:"required,min=7"`
}

func AuthenticationMiddleware(context *gin.Context) {
	header := Header{}
	if err = context.BindHeader(&header); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err = validate.Struct(header); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Token is invalid"})
		return
	}
	token := header.Authorization[7:]
	verifyToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err == nil && verifyToken.Valid {
		context.Next()
	} else {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Token is invalid"})
		return
	}
}
