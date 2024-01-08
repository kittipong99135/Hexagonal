package service

import "hex/repository"

type UserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Phone    string `json:"phone"`
	Rank     string `json:"rank"`
}

type UserResponse struct {
	Email string `json:"email"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Rank  string `json:"rank"`
}

type LoginResponse struct {
	Status        string
	Access_token  string
	Refresh_token string
}

type AuthService interface {
	RegisterService(UserRequest) (*UserResponse, error)
	LoginService(repository.UserAuth) (*LoginResponse, error)
	UserListAll(string) ([]UserResponse, error)
	UserReadById(string) (*UserResponse, error)
}
