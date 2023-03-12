package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/xsicx/highload/internal/usecases/users"
)

type UsersController struct {
	UsersGateway users.Gateway
}

func NewUsersController(usersGateway users.Gateway) *UsersController {
	return &UsersController{
		UsersGateway: usersGateway,
	}
}

func (c *UsersController) login(ctx *fiber.Ctx) error {
	return nil
}

func (c *UsersController) register(ctx *fiber.Ctx) error {
	return nil
}

func (c *UsersController) getByID(ctx *fiber.Ctx) error {
	return nil
}
