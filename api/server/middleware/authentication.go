package middleware

import (
	"github.com/gin-gonic/gin"
)

func (m *middleware) authentication() gin.HandlerFunc {
	return func(c *gin.Context) {}
}
