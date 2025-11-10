package middleware

import "github.com/gin-gonic/gin"

func (m *Middleware) Authorization(c *gin.Context) {
	isAdmin, exists := c.Get("isAdmin")
	if !exists || !isAdmin.(bool) {
		c.JSON(403, gin.H{
			"error": "forbidden: admin access required",
		})
		c.Abort()
		return
	}

	c.Next()
}
