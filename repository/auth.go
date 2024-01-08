package repository

import (
	"gorm.io/gorm"
)

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

type AuthRepository interface {
	Create(User) (*User, error)
	Authentication(UserAuth) (*User, int, error)
	UserExist(string) (int, error)
	UserList(string) ([]User, error)
	UserRead(string) (*User, error)
}
