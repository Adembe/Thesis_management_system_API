package routes

import (
	"go-rest-api/controllers"

	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		login := main.Group("auth")
		{
			login.POST("/login", controllers.Login)
		}
		users := main.Group("users")
		{
			users.GET("/:id", controllers.GetUser)
			users.GET("/", controllers.GetAllUsers)
			users.POST("/", controllers.CreateUser)
			users.DELETE("/:id", controllers.DeleteUser)
			users.PUT("/", controllers.UpdateUser)
		}
		order := main.Group("order")
		{
			order.POST("/", controllers.CreateOrder)
			order.GET("/", controllers.GetAllOrders)
			order.DELETE("/", controllers.DeleteOrders)
		}
	}

	return router
}
