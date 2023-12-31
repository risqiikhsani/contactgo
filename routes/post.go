package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/risqiikhsani/contactgo/controllers"
	"github.com/risqiikhsani/contactgo/middlewares"
)

func SetupPostRoutes(public *gin.RouterGroup) {
	postGroup := public.Group("/posts")
	postGroup.Use(middlewares.AuthMiddleware())
	{
		postGroup.GET("", controllers.GetPosts)
		postGroup.GET("/:post_id", controllers.GetPostById)
		postGroup.PUT("/:post_id", controllers.UpdatePostById)
		postGroup.DELETE("/:post_id", controllers.DeletePostById)
		postGroup.POST("", controllers.CreatePost)
	}
}
