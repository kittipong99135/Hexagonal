package models

import "gorm.io/gorm"

// Regis Struct
type User struct {
	gorm.Model
	Email    string `db:"email"`
	Password string `db:"email"`
	Name     string `db:"name"`
	Age      int    `db:"age"`
	Phone    string `db:"phone"`
	Rank     string `db:"rank"`
	Role     string `db:"role"`
	Status   string `db:"status"`
}

type UserAuth struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Auth Struct
type UserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Phone    string `json:"phone"`
	Rank     string `json:"rank"`
}

type UserResponse struct {
	ID     string `json:"id"`
	Email  string `json:"email"`
	Name   string `json:"name"`
	Phone  string `json:"phone"`
	Rank   string `json:"rank"`
	Status string `json:"status"`
	Role   string `json:"role"`
}

type LoginResponse struct {
	Status        string
	Access_token  string
	Refresh_token string
}
