package repository

import (
	"crud-mvc-go/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository{
	return &UserRepository{DB: db}
}

func (r *UserRepository) GetAllUsers() ([]models.User, error){
	var users []models.User
	err := r.DB.Find(&users).Error
	return users, err
}

func (r *UserRepository) GetUserByID(id_users uint) (models.User, error){
	var user models.User
	err := r.DB.First(&user,"id_users = ?", id_users).Error
	return user, err
}

func (r *UserRepository) CreateUser(user *models.User) error{
	return r.DB.Create(user).Error
}

func (r *UserRepository) UpdateUser(user *models.User) error{
	return r.DB.Model(user).Updates(map[string]interface{}{
		"name": user.Name,
		"email": user.Email,
	}).Error
}

func (r *UserRepository) DeleteUser(id_users uint) error{
	return r.DB.Delete(&models.User{}, "id_users", id_users).Error
}
