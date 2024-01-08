package routes

import (
	"hex/handdler"
	"hex/middleware"
	"hex/repository"
	"hex/service"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Routes(app *fiber.App, db *gorm.DB) {

	api := app.Group("/api")

	// Auth Structure
	authRepository := repository.NewAuthRepository(db)
	authService := service.NewAuthService(authRepository)
	authHandler := handdler.NewAuthHandler(authService)

	// Auth Routes
	auth := api.Group("/auth")
	auth.Post("/register", authHandler.Register)
	auth.Post("/login", authHandler.Login)

	// User Routes
	user := api.Group(
		"/user",
		middleware.RequestAuth(),
		middleware.RefreshAuth(),
	)
	user.Get("/params", authHandler.UserParams)
	user.Get("/list", authHandler.ListAllUser)
	user.Get("/read/:id", authHandler.ReadUserById)
	user.Put("/update/:id", authHandler.ActiveStatus)
	user.Delete("/remove/:id", authHandler.RemoveUser)

}
