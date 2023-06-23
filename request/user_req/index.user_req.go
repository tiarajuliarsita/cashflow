package user_req

import "time"

type UsersReq struct {
	Email       string    `json:"email" binding:"required"`
	Pin         int       `json:"pin" binding:"required"`
	UserName    string    `json:"user_name" binding:"required"`
	PhoneNumber int       `json:"phone_number" binding:"required"`
	BornDate    time.Time `json:"born_date" binding:"required"`
	Saldo       int       `json:"saldo" binding:"required"`
}
