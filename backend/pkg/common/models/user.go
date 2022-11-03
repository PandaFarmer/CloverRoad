package models

import "gorm.io/gorm"

type User struct {
	gorm.Model         // adds ID, created_at etc.
	Email    string `json:"email"`
	UserName string `json:"user_name"`
	FullName    string    `json:"full_name"`
	Password    string    `json:"password"`
	IsSuperUser bool      `json:"is_super_user"`
}
