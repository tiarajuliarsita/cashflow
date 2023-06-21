package models

import "time"

type Users struct {
	ID            int       `json:"id"`
	UserName      string    `json:"user_name"` // Ubah penamaan kolom menjadi user_name
	Email         string    `json:"email"`
	Pin           int    `json:"pin"`
	AccountNumber string    `json:"account_number"`
	PhoneNumber   int       `json:"phone_number"`
	Balance       int       `json:"balance"`
	BornDate      time.Time `json:"born_date"`
}

type Transfer struct {
	ID                     int       `json:"id"`
	UserID                 int       `json:"user_id"`
	Date                   time.Time `json:"date"`
	Amount                 int       `json:"amount"`
	RecipientAccountNumber string    `json:"recipient_account_number"`
}
