package models

import (
	"time"
)

type User struct {
	Id       int    `json:"id" gorm:"primaryKey"`
	Email    string `json:"email"`
	UserName string `json:"user_name"`
	// FirstName   string `json:"first_name"`
	// Surname		string `json:"surname`
	FullName    string    `json:"full_name"`
	Password    string    `json:"password"`
	DateJoined  time.Time `json:"date_joined"`
	IsSuperUser bool      `json:"is_super_user"`
}
