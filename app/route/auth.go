package route

import (
	"keycloak_api_go/app/controller"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.Engine) *gin.RouterGroup {
	authRoutes := router.Group("/auth")
	authRoutes.POST("/register", controller.Register)
	authRoutes.POST("/login", controller.Login)

	return authRoutes
}
