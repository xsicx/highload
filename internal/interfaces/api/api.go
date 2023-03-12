package api

import (
	fiber "github.com/gofiber/fiber/v2"
	"github.com/xsicx/highload/internal/usecases/users"
)

type Dependencies struct {
	UsersGateway users.Gateway
}

func New(router *fiber.App, deps Dependencies) {
	usersController := NewUsersController(deps.UsersGateway)
	router.Post("/login", usersController.login)
	router.Post("/register", usersController.register)
	router.Get("/user/get/:id", usersController.getByID)
}
