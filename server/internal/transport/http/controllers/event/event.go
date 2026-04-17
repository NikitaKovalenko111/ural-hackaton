package event_controller

import (
	"database/sql"
	"strconv"
	"strings"

	event_dto "ural-hackaton/internal/dto/event"
	"ural-hackaton/internal/models"
	auth_service "ural-hackaton/internal/services/handlers/auth"
	event_service "ural-hackaton/internal/services/handlers/event"
	mentor_service "ural-hackaton/internal/services/handlers/mentor"

	"github.com/gofiber/fiber/v2"
)

type EventController struct {
	service       *event_service.EventService
	authService   *auth_service.AuthService
	mentorService *mentor_service.MentorService
}

func Init(service *event_service.EventService, authService *auth_service.AuthService, mentorService *mentor_service.MentorService) *EventController {
	return &EventController{service: service, authService: authService, mentorService: mentorService}
}

func (c *EventController) RegisterRoutes(router fiber.Router) {
	events := router.Group("/events")
	events.Get("/", c.GetAllEvents)
	events.Get("/search", c.SearchEvents)
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

func (c *EventController) SearchEvents(ctx *fiber.Ctx) error {
	if c.service == nil {
		return serviceNotReady("event")
	}

	query := ctx.Query("q")
	if strings.TrimSpace(query) == "" {
		return fiber.NewError(fiber.StatusBadRequest, "q query is required")
	}

	events, err := c.service.SearchEventsByName(query)
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
			return fiber.NewError(fiber.StatusNotFound, "event not found")
		}
		return err
	}

	return ctx.JSON(event)
}

func (c *EventController) CreateEvent(ctx *fiber.Ctx) error {
	if c.service == nil {
		return serviceNotReady("event")
	}
	if c.authService == nil {
		return serviceNotReady("auth")
	}

	sessionToken := ctx.Cookies("session_token")
	if sessionToken == "" {
		return fiber.NewError(fiber.StatusUnauthorized, "unauthorized")
	}

	user, err := c.authService.GetSessionUser(sessionToken)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "unauthorized")
	}

	role := strings.ToLower(strings.TrimSpace(user.Role))
	if role != "mentor" && role != "admin" {
		return fiber.NewError(fiber.StatusForbidden, "only mentors or admins can create events")
	}

	var payload event_dto.CreateEventDto
	if err := ctx.BodyParser(&payload); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid event payload")
	}

	mentorID := payload.MentorId
	if role == "mentor" {
		if c.mentorService == nil {
			return serviceNotReady("mentor")
		}

		mentor, mentorErr := c.mentorService.GetMentorByUserId(user.UserID)
		if mentorErr != nil {
			if mentorErr == sql.ErrNoRows {
				return fiber.NewError(fiber.StatusForbidden, "mentor profile not found")
			}

			return mentorErr
		}

		mentorID = &mentor.MentorId
	}

	if mentorID == nil {
		return fiber.NewError(fiber.StatusBadRequest, "mentor_id is required")
	}

	err = c.service.CreateEvent(payload.Name, payload.Description, payload.StartTime, payload.EndTime, payload.HubId, mentorID)
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
		return fiber.NewError(fiber.StatusBadRequest, "invalid event payload")
	}

	event, err := c.service.UpdateEvent(payload.EventName, payload.Description, payload.StartTime, payload.EndTime, payload.HubId, payload.MentorId, payload.Id)
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
