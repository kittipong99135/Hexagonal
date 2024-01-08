package service

import (
	"context"
	"errors"
	"hex/database"
	"hex/repository"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type authService struct {
	authRepo repository.AuthRepository
}

type AuthenReaponse struct {
	Token string
}

type UserAuth struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewAuthService(authRepo repository.AuthRepository) AuthService {
	return authService{authRepo: authRepo}
}

func (s authService) RegisterService(data UserRequest) (*UserResponse, error) {

	resultList, _ := s.authRepo.UserExist(data.Email)
	if resultList != 0 {
		return nil, errors.New("User Exits")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(data.Password), 10)
	if err != nil {
		return nil, errors.New("Invalid hash passward")
	}

	user := repository.User{
		Email:    data.Email,
		Password: string(hash),
		Name:     data.Name,
		Age:      data.Age,
		Phone:    data.Phone,
		Rank:     data.Rank,
		Role:     "user",
		Status:   "nactive",
	}

	s.authRepo.Create(user)

	response := UserResponse{
		Email: data.Email,
		Name:  data.Name,
		Phone: data.Phone,
		Rank:  data.Rank,
	}
	return &response, nil
}

func (s authService) LoginService(loginReq repository.UserAuth) (*LoginResponse, error) {

	userResult, numRows, _ := s.authRepo.Authentication(loginReq)
	if numRows == 0 {
		return nil, errors.New("Email invalid")
	}

	err := bcrypt.CompareHashAndPassword([]byte(userResult.Password), []byte(loginReq.Password))
	if err != nil {
		return nil, errors.New("Password invalid")
	}

	udid := strconv.Itoa(int(userResult.ID))

	act_token, err := CreateToken(userResult, "JWT_SECRET")
	if err != nil {
		return nil, errors.New("Create accesstoke invalids")
	}
	SetAccessToken("access_token:"+udid, act_token)

	rfh_token, err := CreateToken(userResult, "JWT_REFRESH")
	if err != nil {
		return nil, errors.New("Create accesstoke invalids")
	}
	SetRefreshToken("refresh_token:"+udid, rfh_token)

	responseLogin := LoginResponse{
		Status:        "success",
		Access_token:  GetToken("access_token:" + udid),
		Refresh_token: GetToken("refresh_token:" + udid),
	}
	return &responseLogin, nil
}

func (s authService) UserListAll(uid string) ([]UserResponse, error) {
	users, err := s.authRepo.UserList(uid)
	if err != nil {
		return nil, errors.New("Expexted users")
	}
	userResponses := []UserResponse{}
	for _, user := range users {
		userResponse := UserResponse{
			Email: user.Email,
			Name:  user.Name,
			Phone: user.Phone,
			Rank:  user.Rank,
		}
		userResponses = append(userResponses, userResponse)
	}

	return userResponses, nil
}

func (s authService) UserReadById(uid string) (*UserResponse, error) {
	user, err := s.authRepo.UserRead(uid)
	if err != nil {
		return nil, errors.New("Expexted users")
	}

	userResponse := UserResponse{
		Email: user.Email,
		Name:  user.Name,
		Phone: user.Phone,
		Rank:  user.Rank,
	}

	return &userResponse, nil
}

// Helping Method

func CreateToken(userResult *repository.User, env string) (string, error) {
	cliams := jwt.MapClaims{
		"uid":    userResult.ID,
		"name":   userResult.Name,
		"email":  userResult.Email,
		"role":   userResult.Role,
		"status": userResult.Status,
		"rank":   userResult.Rank,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, cliams)
	return token.SignedString([]byte(os.Getenv("env")))
}

func SetAccessToken(key string, token string) {
	rd := database.RDConn
	ctx := context.Background()
	rd.Set(ctx, key, token, time.Second*10)
}

func SetRefreshToken(key string, token string) {
	rd := database.RDConn
	ctx := context.Background()
	rd.Set(ctx, key, token, 0)
}

func GetToken(key string) string {
	rd := database.RDConn
	ctx := context.Background()
	val, _ := rd.Get(ctx, key).Result()
	return val
}
