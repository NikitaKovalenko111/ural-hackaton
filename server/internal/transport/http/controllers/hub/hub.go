package hub_controller

import (
	"database/sql"
	"strconv"

	hubs_dto "ural-hackaton/internal/dto/hub"
	"ural-hackaton/internal/models"

	"github.com/gofiber/fiber/v2"
)

type HubService interface {
	GetAllHubs() ([]*models.Hub, error)
	GetHubById(id uint64) (*models.Hub, error)
	CreateHub(hub *hubs_dto.CreateHubDto) (*models.Hub, error)
	UpdateHub(hub *models.Hub) (*models.Hub, error)
	DeleteHub(id uint64) error
}

type HubController struct {
	service HubService
}

func Init(service HubService) *HubController {
	return &HubController{service: service}
}

func (c *HubController) RegisterRoutes(router fiber.Router) {
	hubs := router.Group("/hubs")
	hubs.Get("/", c.GetAllHubs)
	hubs.Get("/:id", c.GetHubById)
	hubs.Post("/", c.CreateHub)
	hubs.Put("/", c.UpdateHub)
	hubs.Delete("/:id", c.DeleteHub)
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

	var payload hubs_dto.CreateHubDto
	if err := ctx.BodyParser(&payload); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid hub payload")
	}

	hub, err := c.service.CreateHub(&payload)
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

	hub, err := c.service.UpdateHub(&payload)
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
