package request_controller

import (
	"database/sql"
	"strconv"

	requests_dto "ural-hackaton/internal/dto/request"
	request_service "ural-hackaton/internal/services/handlers/request"

	"github.com/gofiber/fiber/v2"
)

type RequestController struct {
	service *request_service.RequestService
}

func Init(service *request_service.RequestService) *RequestController {
	return &RequestController{service: service}
}

func (c *RequestController) RegisterRoutes(router fiber.Router) {
	router.Post("/requests", c.CreateRequest)
	router.Get("/requests/:id", c.GetRequestById)
	router.Get("/requests/user/:user_id", c.GetRequestsByUserId)
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

func (c *RequestController) CreateRequest(ctx *fiber.Ctx) error {
	if c.service == nil {
		return serviceNotReady("request")
	}

	var payload requests_dto.CreateRequestDto
	if err := ctx.BodyParser(&payload); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid request payload")
	}

	if err := c.service.CreateRequest(payload.Message, payload.UserId); err != nil {
		return err
	}

	return ctx.SendStatus(fiber.StatusCreated)
}

func (c *RequestController) GetRequestById(ctx *fiber.Ctx) error {
	if c.service == nil {
		return serviceNotReady("request")
	}

	id, err := parseUintParam(ctx, "id")
	if err != nil {
		return err
	}

	request, err := c.service.GetRequestById(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return fiber.NewError(fiber.StatusNotFound, "request not found")
		}
		return err
	}

	return ctx.JSON(request)
}

func (c *RequestController) GetRequestsByUserId(ctx *fiber.Ctx) error {
	if c.service == nil {
		return serviceNotReady("request")
	}

	userId, err := parseUintParam(ctx, "user_id")
	if err != nil {
		return err
	}

	requests, err := c.service.GetRequestsByUserId(userId)
	if err != nil {
		if err == sql.ErrNoRows {
			return fiber.NewError(fiber.StatusNotFound, "requests not found")
		}
		return err
	}

	return ctx.JSON(requests)
}
