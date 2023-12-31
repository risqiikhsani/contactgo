package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/risqiikhsani/contactgo/controllers"
	"github.com/risqiikhsani/contactgo/middlewares"
)

func SetupUserRoutes(public *gin.RouterGroup) {
	userGroup := public.Group("/users")
	userGroup.Use(middlewares.AuthMiddleware())
	{
		userGroup.GET("", controllers.GetUsers)
		userGroup.GET("/:id", controllers.GetUserById)
		userGroup.PUT("/:id", controllers.UpdateUserById)
		userGroup.DELETE("/:id", controllers.DeleteUserById)
		userGroup.POST("", controllers.CreateUser)
	}
}
