package service

import "hex/models"

type AuthService interface {
	RegisterService(models.UserRequest) (*models.UserResponse, error)
	LoginService(models.UserAuth) (*models.LoginResponse, error)
	UserListAll(string) ([]models.UserResponse, error)
	UserReadById(string) (*models.UserResponse, error)
	UserRemove(string)
	ActiveStatus(string) (*models.UserResponse, error)
}
