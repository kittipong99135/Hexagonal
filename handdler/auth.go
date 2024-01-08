package handdler

import (
	"fmt"
	"hex/repository"
	"hex/service"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type authHandler struct {
	authSrv service.AuthService
}

func NewAuthHandler(authSrv service.AuthService) authHandler {
	return authHandler{authSrv: authSrv}
}

func (h authHandler) Register(c *fiber.Ctx) error {

	regisReq := service.UserRequest{}
	err := c.BodyParser(&regisReq)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "Invalid",
		})
	}
	resultReq, err := h.authSrv.RegisterService(regisReq)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "Invalid",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  "Success",
		"message": "Success : Register user success.",
		"user":    resultReq,
	})
}

func (h authHandler) Login(c *fiber.Ctx) error {

	loginReq := repository.UserAuth{}
	err := c.BodyParser(&loginReq)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "Invalid",
		})
	}
	resultReq, err := h.authSrv.LoginService(loginReq)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "Invalid",
		})
	}

	return c.Status(200).JSON(resultReq)
}

func (h authHandler) UserParams(c *fiber.Ctx) error {
	type paramsUser struct {
		Uid    string
		Name   string
		Email  string
		Role   string
		Status string
		Rank   string
	}

	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	uid := fmt.Sprintf("%v", claims["uid"])
	name := fmt.Sprintf("%v", claims["name"])
	email := fmt.Sprintf("%v", claims["email"])
	role := fmt.Sprintf("%v", claims["role"])
	status := fmt.Sprintf("%v", claims["status"])
	rank := fmt.Sprintf("%v", claims["rank"])
	params := paramsUser{
		Uid:    uid,
		Name:   name,
		Email:  email,
		Role:   role,
		Status: status,
		Rank:   rank,
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  "Success",
		"message": "Success : Read user success.",
		"params":  params,
	})
}

func (h authHandler) ListAllUser(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	uid := fmt.Sprintf("%v", claims["uid"])
	users, _ := h.authSrv.UserListAll(uid)

	return c.Status(200).JSON(fiber.Map{
		"status":  "Success",
		"message": "Success : List user success.",
		"params":  users,
	})
}

func (h authHandler) ReadUserById(c *fiber.Ctx) error {
	uid := c.Params("id")
	user, err := h.authSrv.UserReadById(uid)
	if err != nil {
		return c.Status(503).JSON(fiber.Map{
			"status":  "error",
			"message": "Error : Refresh token invalid.",
			"error":   err.Error(),
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"status":  "Success",
		"message": "Success : Read user success.",
		"user":    user,
	})
}
