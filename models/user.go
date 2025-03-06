package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	IDUser uint   `gorm:"primaryKey;column:id_users" json:"id_users"`
    Name  string `json:"name"`
    Email string `json:"email"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
    DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

func MigrateUsers(db *gorm.DB){
	db.AutoMigrate(&User{})
}