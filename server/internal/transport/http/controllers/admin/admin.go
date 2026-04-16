package admin_controller

import (
	"database/sql"
	"strconv"

	admin_dto "ural-hackaton/internal/dto/admin"

	"github.com/gofiber/fiber/v2"
)

type AdminService interface {
	GetAllAdmins() ([]*admin_dto.AdminJoinUserDto, error)
	GetAdminById(id uint64) (*admin_dto.AdminJoinUserDto, error)
	CreateAdmin(admin admin_dto.CreateAdminDto) (*admin_dto.AdminJoinUserDto, error)
	DeleteAdmin(id uint64) error
	GetAdminsByFullname(fullname string) ([]*admin_dto.AdminJoinUserDto, error)
}

type AdminController struct {
	service AdminService
}

func Init(service AdminService) *AdminController {
	return &AdminController{service: service}
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

func (c *AdminController) GetAllAdmins(ctx *fiber.Ctx) error {
	if c.service == nil {
		return serviceNotReady("admin")
	}

	admins, err := c.service.GetAllAdmins()
	if err != nil {
		return err
	}

	return ctx.JSON(admins)
}

func (c *AdminController) GetAdminById(ctx *fiber.Ctx) error {
	if c.service == nil {
		return serviceNotReady("admin")
	}

	id, err := parseUintParam(ctx, "id")
	if err != nil {
		return err
	}

	admin, err := c.service.GetAdminById(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return fiber.NewError(fiber.StatusNotFound, "admin not found")
		}
		return err
	}

	return ctx.JSON(admin)
}

func (c *AdminController) CreateAdmin(ctx *fiber.Ctx) error {
	if c.service == nil {
		return serviceNotReady("admin")
	}

	var payload admin_dto.CreateAdminDto
	if err := ctx.BodyParser(&payload); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid admin payload")
	}

	admin, err := c.service.CreateAdmin(payload)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusCreated).JSON(admin)
}

func (c *AdminController) DeleteAdmin(ctx *fiber.Ctx) error {
	if c.service == nil {
		return serviceNotReady("admin")
	}

	id, err := parseUintParam(ctx, "id")
	if err != nil {
		return err
	}

	if err := c.service.DeleteAdmin(id); err != nil {
		return err
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}

func (c *AdminController) GetAdminsByFullname(ctx *fiber.Ctx) error {
	if c.service == nil {
		return serviceNotReady("admin")
	}

	fullname := ctx.Query("fullname")
	if fullname == "" {
		return fiber.NewError(fiber.StatusBadRequest, "fullname query is required")
	}

	admins, err := c.service.GetAdminsByFullname(fullname)
	if err != nil {
		if err == sql.ErrNoRows {
			return fiber.NewError(fiber.StatusNotFound, "admins not found")
		}
		return err
	}

	return ctx.JSON(admins)
}
