package controllers

import (
	"crud-mvc-go/models"
	"crud-mvc-go/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserRepo *repository.UserRepository
}

func NewUserController(repo *repository.UserRepository) *UserController{
	return &UserController{UserRepo: repo}
}

func (c *UserController) GetAllUsers(ctx *gin.Context) {
	users, err := c.UserRepo.GetAllUsers()
	if err != nil{
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()}) 
		return
	}
	ctx.JSON(http.StatusOK, users)
}

func (c *UserController) GetUserByID(ctx *gin.Context){
	id_user, err := strconv.Atoi(ctx.Param("id_user"))
	if err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	user, err := c.UserRepo.GetUserByID(uint(id_user))
	if err != nil{
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User Not Found"})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (c *UserController) CreateUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.UserRepo.CreateUser(&user); err != nil{
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, user)
}

func (c *UserController) UpdateUser(ctx *gin.Context) {
	id_user, err := strconv.Atoi(ctx.Param("id_user"))
	if err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user.IDUser = uint(id_user)

	if err := c.UserRepo.UpdateUser(&user); err != nil{
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (c *UserController) DeleteUser(ctx *gin.Context){
	id_user, err := strconv.Atoi(ctx.Param("id_user"))
	if err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := c.UserRepo.DeleteUser(uint(id_user)); err != nil{
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User deletes successfully"})
}