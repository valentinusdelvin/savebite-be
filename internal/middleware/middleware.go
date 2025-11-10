package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/valentinusdelvin/savebite-be/internal/pkg/jwt"
)

type MiddlewareItf interface {
	Authentication(c *gin.Context)
	Authorization(c *gin.Context)
}

type Middleware struct {
	jwt.JWTItf
}

func NewMiddleware(jwt jwt.JWTItf) MiddlewareItf {
	return &Middleware{
		JWTItf: jwt,
	}
}
