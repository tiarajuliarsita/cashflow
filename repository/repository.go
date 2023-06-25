package repository

import (
	"log"
	"mini_Atm/models"
	"gorm.io/gorm"
)

type Repository interface {
	FindByID(ID int) (models.Users, error)
	FindAll() ([]models.Users, error)
}

type repository struct {
	DB *gorm.DB
}

func NewRepository(DB *gorm.DB) *repository {
	return &repository{DB}
}

func (r *repository) FindAll() ([]models.Users, error) {
	var users []models.Users
	err := r.DB.Table("users").Find(&users).Error
	if err != nil {
		log.Println(err)
	}
	return users, err
}

func (r *repository) FindByID(ID int) (models.Users, error) {
	var user models.Users
	err := r.DB.Table("users").Find(&user, ID).Error
	if err != nil {
		log.Println(err)
	}
	return user, err
}

// func (r *repository) Create(user_req.UsersReq) (models.Users, error) {
// 	var user models.Users
// 	err := r.DB.Table("users").Create(&user).Error
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	return user, err
// }
