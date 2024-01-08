package repository

import (
	"hex/models"

	_ "gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return authRepository{db: db}
}

func (r authRepository) Create(data models.User) (*models.User, error) {
	result := r.db.Create(&data)
	return &data, result.Error
}

func (r authRepository) Authentication(targetEmail models.UserAuth) (*models.User, int, error) {
	var resultUser models.User
	result := r.db.Find(&resultUser, "email =?", targetEmail.Email)
	if result.Error != nil {
		return nil, 0, result.Error
	}

	return &resultUser, int(result.RowsAffected), result.Error
}

func (r authRepository) UserExist(targetEmail string) (int, error) {
	var userExist models.User
	result := r.db.Find(&userExist, "email =?", targetEmail)
	if result.Error != nil {
		return 0, result.Error
	}
	return int(result.RowsAffected), nil
}

func (r authRepository) UserList(uid string) ([]models.User, error) {
	var userList []models.User
	result := r.db.Find(&userList, "id !=?", uid)
	if result.Error != nil {
		return nil, result.Error
	}
	return userList, nil
}

func (r authRepository) UserRead(uid string) (*models.User, error) {
	var userList models.User
	result := r.db.Find(&userList, "id =?", uid)
	if result.Error != nil {
		return nil, result.Error
	}
	return &userList, nil
}

func (r authRepository) UserRemove(uid string) {
	var userRemove *models.User
	r.db.Delete(&userRemove, uid)
}

func (r authRepository) UserStatus(uid string) (*models.User, int, error) {
	var user models.User

	result := r.db.First(&user, "id =?", uid)
	if result.Error != nil {
		return nil, 0, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, 0, result.Error
	}

	var activeUser models.User
	if user.Status != "active" {
		activeUser = models.User{Status: "active"}
	} else {
		activeUser = models.User{Status: "nactive"}
	}

	r.db.Where("id = ?", uid).Updates(&activeUser)

	return &activeUser, 1, nil
}
