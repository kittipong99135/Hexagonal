package repository

import (
	_ "gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return authRepository{db: db}
}

func (r authRepository) Create(data User) (*User, error) {
	result := r.db.Create(&data)
	return &data, result.Error
}

func (r authRepository) Authentication(targetEmail UserAuth) (*User, int, error) {
	var resultUser User
	result := r.db.Find(&resultUser, "email =?", targetEmail.Email)
	if result.Error != nil {
		return nil, 0, result.Error
	}

	return &resultUser, int(result.RowsAffected), result.Error
}

func (r authRepository) UserExist(targetEmail string) (int, error) {
	var userExist User
	result := r.db.Find(&userExist, "email =?", targetEmail)
	if result.Error != nil {
		return 0, result.Error
	}
	return int(result.RowsAffected), nil
}

func (r authRepository) UserList(uid string) ([]User, error) {
	var userList []User
	result := r.db.Find(&userList, "id !=?", uid)
	if result.Error != nil {
		return nil, result.Error
	}
	return userList, nil
}

func (r authRepository) UserRead(uid string) (*User, error) {
	var userList User
	result := r.db.Find(&userList, "id =?", uid)
	if result.Error != nil {
		return nil, result.Error
	}
	return &userList, nil
}
