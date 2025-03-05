package routes

import (
	"crud-mvc-go/controllers"

	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(router *gin.Engine, userController *controllers.UserController){
	userGroup := router.Group("/users")
	{
		userGroup.GET("", userController.GetAllUsers)
		userGroup.GET("/:id_user", userController.GetUserByID)
        userGroup.POST("", userController.CreateUser)
        userGroup.PUT("/:id_user", userController.UpdateUser)
        userGroup.DELETE("/:id_user", userController.DeleteUser)
	}
}