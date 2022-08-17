package models

type Model3D struct {
	Id          int    `json:"id" gorm:"primaryKey"`
	Email		string `json:"email"`
	UserName	string `json:"user_name"`
	FirstName   string `json:"first_name"`
	LastName	string `json:"last_name`
	DateJoined  string `json:"date_joined"`
}