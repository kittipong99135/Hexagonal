package repository

import (
	"hex/models"
)

var User models.User

type AuthRepository interface {
	Create(models.User) (*models.User, error)
	Authentication(models.UserAuth) (*models.User, int, error)
	UserExist(string) (int, error)
	UserList(string) ([]models.User, error)
	UserRead(string) (*models.User, error)
	UserRemove(string)
	UserStatus(string) (*models.User, int, error)
}
