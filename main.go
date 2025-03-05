package main

import (
	"crud-mvc-go/config"
	"crud-mvc-go/controllers"
	"crud-mvc-go/models"
	"crud-mvc-go/repository"
	"crud-mvc-go/routes"

	"github.com/gin-gonic/gin"
)


func main() {
	config.ConnectDB()
	models.MigrateUsers(config.DB)

	userRepo := repository.NewUserRepository(config.DB)
	userController := controllers.NewUserController(userRepo)

	r := gin.Default()
	
	routes.SetupUserRoutes(r, userController)

	r.Run(":8080")
}