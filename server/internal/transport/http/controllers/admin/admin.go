package admin_controller

import (
	"database/sql"
	"strconv"
	admin_dto "ural-hackaton/internal/dto/admin"
	admin_service "ural-hackaton/internal/services/handlers/admin"

	"github.com/gofiber/fiber/v2"
)

type AdminController struct {
	service *admin_service.AdminService
}

func Init(service *admin_service.AdminService) *AdminController {
	return &AdminController{service: service}
}

func (c *AdminController) RegisterRoutes(router fiber.Router) {
	admins := router.Group("/admins")
	admins.Get("/", c.GetAllAdmins)
	admins.Get("/search", c.GetAdminsByFullname)
	admins.Get("/:id", c.GetAdminById)
	admins.Post("/", c.CreateAdmin)
	admins.Delete("/:id", c.DeleteAdmin)
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
	admins, err := c.service.GetAllAdmins()
	if err != nil {
		return err
	}

	return ctx.JSON(admins)
}

func (c *AdminController) GetAdminById(ctx *fiber.Ctx) error {
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
	var payload admin_dto.CreateAdminDto
	if err := ctx.BodyParser(&payload); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid admin payload")
	}

	admin, err := c.service.CreateAdmin(payload.UserId)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusCreated).JSON(admin)
}

func (c *AdminController) DeleteAdmin(ctx *fiber.Ctx) error {
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
	fullname := ctx.Query("fullname")
	if fullname == "" {
		return fiber.NewError(fiber.StatusBadRequest, "fullname query is required")
	}

	admins, err := c.service.GetAdminByFullname(fullname)
	if err != nil {
		if err == sql.ErrNoRows {
			return fiber.NewError(fiber.StatusNotFound, "admins not found")
		}
		return err
	}

	return ctx.JSON(admins)
}
