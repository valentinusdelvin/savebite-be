package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (m *Middleware) Authentication(c *gin.Context) {
	header := c.GetHeader("Authorization")
	if header == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "token is required",
		})
		c.Abort()
		return
	}

	token := strings.Split(header, " ")[1]
	userId, IsAdmin, err := m.JWTItf.ValidateToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "invalid token/Failed to validate token",
		})
		c.Abort()
		return
	}

	c.Set("userID", userId)
	c.Set("isAdmin", IsAdmin)

	c.Next()
}
