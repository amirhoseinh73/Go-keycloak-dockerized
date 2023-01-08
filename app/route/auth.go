package route

import (
	"keycloak_api_go/app/controller"
	"keycloak_api_go/app/middleware"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.Engine) *gin.RouterGroup {
	routes := router.Group("/auth")
	routes.POST("/register", controller.Register)
	routes.POST("/login", controller.Login)

	protectedRoutes := router.Group("/user")
	protectedRoutes.Use(middleware.MiddlewareAuthKC())
	protectedRoutes.GET("/info", controller.Info)

	return routes
}
