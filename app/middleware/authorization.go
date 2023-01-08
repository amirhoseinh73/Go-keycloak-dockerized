package middleware

import (
	"github.com/gin-gonic/gin"
)

func MiddlewareAuthKC() gin.HandlerFunc {
	return func(context *gin.Context) {
		// 	context.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication required!"})
		// 	context.Abort()
		// 	return
		context.Next()
	}
}
