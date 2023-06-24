package service

import (
	"log"
	"mini_Atm/models"
	"mini_Atm/repository"
)

type Service interface {
	FindByID(ID int) (models.Users, error)
	FindAll() ([]models.Users, error)
	// Create(user_req.UsersReq) (models.Users, error)
}

type service struct {
	repository repository.Repository
}

func NewService(repository repository.Repository) *service {
	return &service{repository}
}
func (s *service) FindAll() ([]models.Users, error) {
	users, err := s.repository.FindAll()
	return users, err
}

func (s *service) FindByID(ID int) (models.Users, error) {
	var user models.Users
	user, err := s.repository.FindByID(ID)
	if err != nil {
		log.Println(err)
	}
	return user, err
}

// func (s *service) Create(user_req.UsersReq) (models.Users, error) {
// 	var userReq  user_req.UsersReq
// 	userReq := user_req.UsersReq{
// 		Email:         userReq.Email,
// 		Pin:           userReq.Pin,
// 		PhoneNumber:   userReq.PhoneNumber,
// 		UserName:      userReq.UserName,
// 		BornDate:      userReq.BornDate,
// 		Saldo:         userReq.Saldo,
// 		AccountNumber: user_controllers.GetNumberAccount(7),
// 	}
// 	newUser, err := s.repository.Create(userReq)
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	return newUser, err
// }
