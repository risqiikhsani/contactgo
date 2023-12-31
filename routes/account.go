package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/risqiikhsani/contactgo/controllers"
)

func SetupAccountRoutes(public *gin.RouterGroup) {
	public.POST("/register", controllers.Register)
	public.POST("/login", controllers.Login)
}
