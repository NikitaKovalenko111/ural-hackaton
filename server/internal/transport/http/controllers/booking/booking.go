package booking_controller

import (
	"database/sql"
	"strconv"

	booking_dto "ural-hackaton/internal/dto/booking"
	"ural-hackaton/internal/models"
	booking_service "ural-hackaton/internal/services/handlers/booking"

	"github.com/gofiber/fiber/v2"
)

type BookingController struct {
	service *booking_service.BookingService
}

func Init(service *booking_service.BookingService) *BookingController {
	return &BookingController{service: service}
}

func (c *BookingController) RegisterRoutes(router fiber.Router) {
	bookings := router.Group("/bookings")
	bookings.Get("/", c.GetAllBookings)
	bookings.Get("/:id", c.GetBookingById)
	bookings.Get("/user/:user_id", c.GetBookingsByUserId)
	bookings.Post("/", c.CreateBooking)
	bookings.Put("/", c.UpdateBooking)
	bookings.Delete("/:id", c.DeleteBooking)
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

func (c *BookingController) CreateBooking(ctx *fiber.Ctx) error {
	if c.service == nil {
		return serviceNotReady("booking")
	}

	var payload booking_dto.CreateBookingDto
	if err := ctx.BodyParser(&payload); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid booking payload")
	}

	booking, err := c.service.CreateBooking(payload.BookingDate, payload.BookingZone, payload.BookingSlots, payload.UserId)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusCreated).JSON(booking)
}

func (c *BookingController) GetBookingById(ctx *fiber.Ctx) error {
	if c.service == nil {
		return serviceNotReady("booking")
	}

	id, err := parseUintParam(ctx, "id")
	if err != nil {
		return err
	}

	booking, err := c.service.GetBookingById(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return fiber.NewError(fiber.StatusNotFound, "booking not found")
		}
		return err
	}

	return ctx.JSON(booking)
}

func (c *BookingController) GetBookingsByUserId(ctx *fiber.Ctx) error {
	if c.service == nil {
		return serviceNotReady("booking")
	}

	userId, err := parseUintParam(ctx, "user_id")
	if err != nil {
		return err
	}

	bookings, err := c.service.GetBookingsByUserId(userId)
	if err != nil {
		if err == sql.ErrNoRows {
			return fiber.NewError(fiber.StatusNotFound, "bookings not found")
		}
		return err
	}

	return ctx.JSON(bookings)
}

func (c *BookingController) GetAllBookings(ctx *fiber.Ctx) error {
	if c.service == nil {
		return serviceNotReady("booking")
	}

	bookings, err := c.service.GetAllBookings()
	if err != nil {
		return err
	}

	return ctx.JSON(bookings)
}

func (c *BookingController) UpdateBooking(ctx *fiber.Ctx) error {
	if c.service == nil {
		return serviceNotReady("booking")
	}

	var payload models.Booking
	if err := ctx.BodyParser(&payload); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid booking payload")
	}

	booking, err := c.service.UpdateBooking(payload.BookingDate, payload.BookingZone, payload.BookingSlots, payload.UserId, payload.Id)
	if err != nil {
		return err
	}

	return ctx.JSON(booking)
}

func (c *BookingController) DeleteBooking(ctx *fiber.Ctx) error {
	if c.service == nil {
		return serviceNotReady("booking")
	}

	id, err := parseUintParam(ctx, "id")
	if err != nil {
		return err
	}

	if err := c.service.DeleteBooking(id); err != nil {
		return err
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}
