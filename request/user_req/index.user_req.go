package user_req

import "time"

type UsersReq struct {
	Email       string    `json:"email" binding:"required"`
	Pin         string    `json:"pin" binding:"required"`
	UserName    string    `json:"user_name" binding:"required"`
	PhoneNumber int       `json:"phone_number" binding:"required"`
	BornDate    time.Time `json:"born_date" binding:"required"`
	Balance     int       `json:"balance" binding:"required"`
}
