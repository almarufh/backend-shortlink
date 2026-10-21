package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (m *Middleware) Cors(c *gin.Context) {
	for _, o := range m.allowedOrigins {
		if c.GetHeader("Origin") == o {
			c.Header("Access-Control-Allow-Origin", o)
		}
	}

	c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
	c.Header("Access-Control-Allow-Methods", "DELETE, OPTIONS, PATCH")
	if c.Request.Method == http.MethodOptions {
		c.AbortWithStatus(http.StatusNoContent)
		return
	}
	c.Next()
}
