package user_controller

import (
	"database/sql"
	"strconv"

	user_dto "ural-hackaton/internal/dto/user"
	user_service "ural-hackaton/internal/services/handlers/user"

	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	service *user_service.UserService
}

func Init(service *user_service.UserService) *UserController {
	return &UserController{service: service}
}

func serviceNotReady(entity string) error {
	return fiber.NewError(fiber.StatusNotImplemented, entity+" service is not wired yet")
}

func parseUintParam(c *fiber.Ctx, key string) (uint64, error) {
	value := c.Params(key)
	parsed, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		return 0, fiber.NewError(fiber.StatusBadRequest, "invalid "+key)
	}

	return parsed, nil
}

func (c *UserController) CreateUser(ctx *fiber.Ctx) error {
	if c.service == nil {
		return serviceNotReady("user")
	}

	var payload user_dto.CreateUserDto
	if err := ctx.BodyParser(&payload); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid user payload")
	}

	if err := c.service.CreateUser(payload.Fullname, payload.Role); err != nil {
		return err
	}

	return ctx.SendStatus(fiber.StatusCreated)
}

func (c *UserController) GetUserById(ctx *fiber.Ctx) error {
	if c.service == nil {
		return serviceNotReady("user")
	}

	id, err := parseUintParam(ctx, "id")
	if err != nil {
		return err
	}

	user, err := c.service.GetUserById(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return fiber.NewError(fiber.StatusNotFound, "user not found")
		}
		return err
	}

	return ctx.JSON(user)
}

func (c *UserController) GetUserByFullname(ctx *fiber.Ctx) error {
	if c.service == nil {
		return serviceNotReady("user")
	}

	fullname := ctx.Query("fullname")
	if fullname == "" {
		return fiber.NewError(fiber.StatusBadRequest, "fullname query is required")
	}

	user, err := c.service.GetUserByFullname(fullname)
	if err != nil {
		if err == sql.ErrNoRows {
			return fiber.NewError(fiber.StatusNotFound, "user not found")
		}
		return err
	}

	return ctx.JSON(user)
}

func (c *UserController) GetUsersByRole(ctx *fiber.Ctx) error {
	if c.service == nil {
		return serviceNotReady("user")
	}

	role := ctx.Query("role")
	if role == "" {
		return fiber.NewError(fiber.StatusBadRequest, "role query is required")
	}

	users, err := c.service.GetUsersByRole(role)
	if err != nil {
		if err == sql.ErrNoRows {
			return fiber.NewError(fiber.StatusNotFound, "users not found")
		}
		return err
	}

	return ctx.JSON(users)
}
