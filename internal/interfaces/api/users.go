package api

import (
	"math"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"github.com/xsicx/highload/internal/usecases/dto"
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
	request := dto.LoginUserRequest{}
	if err := ctx.BodyParser(&request); err != nil {
		return errors.Wrap(err, "body parser error")
	}

	manager := users.NewUsersManager(c.UsersGateway)

	userToken, err := manager.Login(ctx.UserContext(), request)
	if err != nil {
		return err
	}

	response := dto.UserTokenResponse{
		Token: userToken,
	}

	ctx.Status(fiber.StatusOK)

	return ctx.JSON(response)
}

func (c *UsersController) register(ctx *fiber.Ctx) error {
	request := dto.RegisterUserRequest{}
	if err := ctx.BodyParser(&request); err != nil {
		return errors.Wrap(err, "body parser error")
	}

	manager := users.NewUsersManager(c.UsersGateway)

	user, err := manager.Register(ctx.UserContext(), request)
	if err != nil {
		return err
	}

	response := dto.UserIDResponse{UserID: user.ID.String()}

	ctx.Status(fiber.StatusCreated)

	return ctx.JSON(response)
}

func (c *UsersController) getByID(ctx *fiber.Ctx) error {
	userID, err := uuid.FromString(ctx.Params("id"))
	if err != nil {
		return errors.Wrap(err, "uuid parser error")
	}

	manager := users.NewUsersManager(c.UsersGateway)

	user, err := manager.GetByID(ctx.UserContext(), userID)
	if err != nil {
		return err
	}

	if user == nil {
		ctx.Status(fiber.StatusNotFound)

		return nil
	}

	age := math.Floor(time.Now().Sub(user.Birthdate).Hours() / 24 / 365)

	userDTO := dto.UserResponse{
		ID:         user.ID.String(),
		FirstName:  user.FirstName,
		SecondName: user.SecondName,
		Age:        int(age),
		Birthdate:  user.Birthdate.Format("2006-01-02"),
		Biography:  user.Biography,
		City:       user.City,
	}

	ctx.Status(fiber.StatusOK)

	return ctx.JSON(userDTO)
}
