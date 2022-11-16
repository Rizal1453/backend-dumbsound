package models

import "time"

type User struct {
	ID          int           `json:"id"`
	Email       string        `json:"email"`
	Password    string        `json:"password"`
	FullName    string        `json:"fullname"`
	Gender      string        `json:"gender"`
	Phone       string        `json:"phone"`
	Address     string        `json:"address"`
	Role        string        `json:"role"`
	Image       string        `json:"image"`
	Status      string        `json:"status"`
	Transaction []Transaction `json:"transaction"`
	CreatedAt   time.Time `json:"-"`
	UpdateAt time.Time `json:"-"`
}
type UserProfile struct{
	ID int `json:"id"`
	FullName string `json:"fullname"`
	Email string `json:"email"`
	Image string `json:"image"`
	Gender string `json:"gender"`
	Phone string `json:"phone"`
	Address string  `json:"address"` 
}
func (UserProfile) TableName() string{
	return "users"
}
