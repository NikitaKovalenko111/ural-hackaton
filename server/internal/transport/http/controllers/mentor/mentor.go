package mentor_controller

import (
	"database/sql"
	"strconv"

	mentor_dto "ural-hackaton/internal/dto/mentor"
	mentor_service "ural-hackaton/internal/services/handlers/mentor"

	"github.com/gofiber/fiber/v2"
)

type MentorController struct {
	service *mentor_service.MentorService
}

func Init(service *mentor_service.MentorService) *MentorController {
	return &MentorController{service: service}
}

func (c *MentorController) RegisterRoutes(router fiber.Router) {
	mentors := router.Group("/mentors")
	mentors.Get("/", c.GetAllMentors)
	mentors.Post("/", c.CreateMentor)
	mentors.Get("/search", c.GetMentorsByFullname)
	mentors.Get("/user/:user_id", c.GetMentorByUserId)
	mentors.Get("/:id", c.GetMentorById)

}

func (c *MentorController) GetMentorByUserId(ctx *fiber.Ctx) error {
	if c.service == nil {
		return serviceNotReady("mentor")
	}

	userId, err := parseUintParam(ctx, "user_id")
	if err != nil {
		return err
	}

	mentor, err := c.service.GetMentorByUserId(userId)
	if err != nil {
		if err == sql.ErrNoRows {
			return fiber.NewError(fiber.StatusNotFound, "mentor not found")
		}
		return err
	}

	return ctx.JSON(mentor)
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

func (c *MentorController) CreateMentor(ctx *fiber.Ctx) error {
	if c.service == nil {
		return serviceNotReady("mentor")
	}

	var payload mentor_dto.CreateMentorDto
	if err := ctx.BodyParser(&payload); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid mentor payload")
	}

	mentor, err := c.service.CreateMentor(payload.UserId, payload.HubId)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusCreated).JSON(mentor)
}

func (c *MentorController) GetAllMentors(ctx *fiber.Ctx) error {
	if c.service == nil {
		return serviceNotReady("mentor")
	}

	mentors, err := c.service.GetAllMentors()
	if err != nil {
		return err
	}

	return ctx.JSON(mentors)
}

func (c *MentorController) GetMentorById(ctx *fiber.Ctx) error {
	if c.service == nil {
		return serviceNotReady("mentor")
	}

	id, err := parseUintParam(ctx, "id")
	if err != nil {
		return err
	}

	mentor, err := c.service.GetMentorById(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return fiber.NewError(fiber.StatusNotFound, "mentor not found")
		}
		return err
	}

	return ctx.JSON(mentor)
}

func (c *MentorController) GetMentorsByFullname(ctx *fiber.Ctx) error {
	if c.service == nil {
		return serviceNotReady("mentor")
	}

	fullname := ctx.Query("fullname")
	if fullname == "" {
		return fiber.NewError(fiber.StatusBadRequest, "fullname query is required")
	}

	mentors, err := c.service.GetMentorByFullname(fullname)
	if err != nil {
		if err == sql.ErrNoRows {
			return fiber.NewError(fiber.StatusNotFound, "mentors not found")
		}
		return err
	}

	return ctx.JSON(mentors)
}
