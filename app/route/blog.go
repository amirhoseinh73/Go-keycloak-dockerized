package route

import (
	"keycloak_api_go/app/middleware"

	"github.com/gin-gonic/gin"
)

func BlogRoutes(router *gin.Engine) *gin.RouterGroup {
	blogRoutes := router.Group("/api/blog")
	blogRoutes.Use(middleware.MiddlewareAuthKC())
	// blogRoutes.POST("/", controller.AddBlog)
	// blogRoutes.GET("/", controller.GetAllBlogs)

	return blogRoutes
}
