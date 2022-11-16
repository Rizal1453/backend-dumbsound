package models

import "time"

type Transaction struct {
	ID         int         `json:"id"`
	UserID     int         `json:"-"`
	User       UserProfile `json:"user"`
	Total      int         `json:"total"`
	Status     string      `json:"status"`
	Limit      int         `json:"limit"`
	StatusUser string      `json:"status_user`
	CreatedAt  time.Time `json:"create_at"`
	UpdateAt time.Time `json:"-"`
}