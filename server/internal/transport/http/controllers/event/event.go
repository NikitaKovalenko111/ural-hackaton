package event_controller

import (
	"database/sql"
	"strconv"

	event_dto "ural-hackaton/internal/dto/event"
	"ural-hackaton/internal/models"
	event_service "ural-hackaton/internal/services/handlers/event"

	"github.com/gofiber/fiber/v2"
)

type EventController struct {
	service *event_service.EventService
}

func Init(service *event_service.EventService) *EventController {
	return &EventController{service: service}
}

func (c *EventController) RegisterRoutes(router fiber.Router) {
	events := router.Group("/events")
	events.Get("/", c.GetAllEvents)
	events.Get("/:id", c.GetEventById)
	events.Post("/", c.CreateEvent)
	events.Put("/", c.UpdateEvent)
	events.Delete("/:id", c.DeleteEvent)
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

func (c *EventController) GetAllEvents(ctx *fiber.Ctx) error {
	if c.service == nil {
		return serviceNotReady("event")
	}

	events, err := c.service.GetAllEvents()
	if err != nil {
		return err
	}

	return ctx.JSON(events)
}

func (c *EventController) GetEventById(ctx *fiber.Ctx) error {
	if c.service == nil {
		return serviceNotReady("event")
	}

	id, err := parseUintParam(ctx, "id")
	if err != nil {
		return err
	}

	event, err := c.service.GetEventById(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return fiber.NewError(fiber.StatusNotFound, "hub not found")
		}
		return err
	}

	return ctx.JSON(event)
}

func (c *EventController) CreateEvent(ctx *fiber.Ctx) error {
	if c.service == nil {
		return serviceNotReady("event")
	}

	var payload event_dto.CreateEventDto
	if err := ctx.BodyParser(&payload); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid hub payload")
	}

	err := c.service.CreateEvent(payload.Name, payload.Description, payload.StartTime, payload.EndTime, payload.HubId)
	if err != nil {
		return err
	}

	return ctx.SendStatus(fiber.StatusCreated)
}

func (c *EventController) UpdateEvent(ctx *fiber.Ctx) error {
	if c.service == nil {
		return serviceNotReady("event")
	}

	var payload models.Event
	if err := ctx.BodyParser(&payload); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid hub payload")
	}

	event, err := c.service.UpdateEvent(payload.EventName, payload.Description, payload.StartTime, payload.EndTime, payload.HubId, payload.Id)
	if err != nil {
		return err
	}

	return ctx.JSON(event)
}

func (c *EventController) DeleteEvent(ctx *fiber.Ctx) error {
	if c.service == nil {
		return serviceNotReady("event")
	}

	id, err := parseUintParam(ctx, "id")
	if err != nil {
		return err
	}

	if err := c.service.DeleteEvent(id); err != nil {
		return err
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}
