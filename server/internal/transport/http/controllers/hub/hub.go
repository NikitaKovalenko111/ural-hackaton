package hub_controller

import (
	"database/sql"
	"strconv"

	hub_dto "ural-hackaton/internal/dto/hub"
	"ural-hackaton/internal/models"
	hub_service "ural-hackaton/internal/services/handlers/hub"

	"github.com/gofiber/fiber/v2"
)

type HubController struct {
	service *hub_service.HubService
}

func Init(service *hub_service.HubService) *HubController {
	return &HubController{service: service}
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

func (c *HubController) GetAllHubs(ctx *fiber.Ctx) error {
	if c.service == nil {
		return serviceNotReady("hub")
	}

	hubs, err := c.service.GetAllHubs()
	if err != nil {
		return err
	}

	return ctx.JSON(hubs)
}

func (c *HubController) GetHubById(ctx *fiber.Ctx) error {
	if c.service == nil {
		return serviceNotReady("hub")
	}

	id, err := parseUintParam(ctx, "id")
	if err != nil {
		return err
	}

	hub, err := c.service.GetHubById(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return fiber.NewError(fiber.StatusNotFound, "hub not found")
		}
		return err
	}

	return ctx.JSON(hub)
}

func (c *HubController) CreateHub(ctx *fiber.Ctx) error {
	if c.service == nil {
		return serviceNotReady("hub")
	}

	var payload hub_dto.CreateHubDto
	if err := ctx.BodyParser(&payload); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid hub payload")
	}

	hub, err := c.service.CreateHub(payload.Name)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusCreated).JSON(hub)
}

func (c *HubController) UpdateHub(ctx *fiber.Ctx) error {
	if c.service == nil {
		return serviceNotReady("hub")
	}

	var payload models.Hub
	if err := ctx.BodyParser(&payload); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid hub payload")
	}

	hub, err := c.service.UpdateHub(payload.HubName, payload.Id)
	if err != nil {
		return err
	}

	return ctx.JSON(hub)
}

func (c *HubController) DeleteHub(ctx *fiber.Ctx) error {
	if c.service == nil {
		return serviceNotReady("hub")
	}

	id, err := parseUintParam(ctx, "id")
	if err != nil {
		return err
	}

	if err := c.service.DeleteHub(id); err != nil {
		return err
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}
