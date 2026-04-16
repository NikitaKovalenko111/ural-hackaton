package mentor_controller

import (
	"database/sql"
	"strconv"

	mentor_dto "ural-hackaton/internal/dto/mentor"

	"github.com/gofiber/fiber/v2"
)

type MentorService interface {
	CreateMentor(mentor mentor_dto.CreateMentorDto) (*mentor_dto.MentorJoinUserDto, error)
	GetMentorById(id uint64) (*mentor_dto.MentorJoinUserDto, error)
	GetMentorsByFullname(fullname string) ([]*mentor_dto.MentorJoinUserDto, error)
	GetMentorsByRole(role string) ([]*mentor_dto.MentorJoinUserDto, error)
}

type MentorController struct {
	service MentorService
}

func Init(service MentorService) *MentorController {
	return &MentorController{service: service}
}

func (c *MentorController) RegisterRoutes(router fiber.Router) {
	router.Post("/mentors", c.CreateMentor)
	router.Get("/mentors/:id", c.GetMentorById)
	router.Get("/mentors/search", c.GetMentorsByFullname)
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

	mentor, err := c.service.CreateMentor(payload)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusCreated).JSON(mentor)
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

	mentors, err := c.service.GetMentorsByFullname(fullname)
	if err != nil {
		if err == sql.ErrNoRows {
			return fiber.NewError(fiber.StatusNotFound, "mentors not found")
		}
		return err
	}

	return ctx.JSON(mentors)
}

func (c *MentorController) GetMentorsByRole(ctx *fiber.Ctx) error {
	if c.service == nil {
		return serviceNotReady("mentor")
	}

	role := ctx.Query("role")
	if role == "" {
		return fiber.NewError(fiber.StatusBadRequest, "role query is required")
	}

	mentors, err := c.service.GetMentorsByRole(role)
	if err != nil {
		if err == sql.ErrNoRows {
			return fiber.NewError(fiber.StatusNotFound, "mentors not found")
		}
		return err
	}

	return ctx.JSON(mentors)
}
