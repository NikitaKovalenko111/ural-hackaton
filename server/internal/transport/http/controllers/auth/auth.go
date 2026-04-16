package auth_controller

import (
	"database/sql"

	auth_dto "ural-hackaton/internal/dto/auth"
	auth_service "ural-hackaton/internal/services/handlers/auth"

	"github.com/gofiber/fiber/v2"
)

type AuthController struct {
	service *auth_service.AuthService
}

func Init(service *auth_service.AuthService) *AuthController {
	return &AuthController{service: service}
}

func (c *AuthController) RegisterRoutes(router fiber.Router) {
	auth := router.Group("/auth")

	// Запрос магической ссылки (ввод email)
	auth.Post("/request", c.RequestMagicLink)

	// Верификация ссылки (переход по токену)
	// GET, так как пользователь переходит по ссылке из письма
	auth.Get("/verify", c.VerifyMagicLink)
}

func serviceNotReady(entity string) error {
	return fiber.NewError(fiber.StatusNotImplemented, entity+" service is not wired yet")
}

// RequestMagicLink: POST /auth/request
// Body: {"email": "user@example.com"}
func (c *AuthController) RequestMagicLink(ctx *fiber.Ctx) error {
	if c.service == nil {
		return serviceNotReady("auth")
	}

	var payload auth_dto.RequestMagicLinkDto
	if err := ctx.BodyParser(&payload); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid request payload")
	}

	// Вызываем сервис.
	// Важно: даже если пользователя нет, мы возвращаем 200, чтобы не светить базу (защита от enum)
	if err := c.service.RequestMagicLink(payload.Email); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "failed to send magic link")
	}

	// Всегда одинаковый ответ для безопасности
	return ctx.JSON(fiber.Map{
		"message": "Если email зарегистрирован, ссылка отправлена",
	})
}

// VerifyMagicLink: GET /auth/verify?token=xxx
func (c *AuthController) VerifyMagicLink(ctx *fiber.Ctx) error {
	if c.service == nil {
		return serviceNotReady("auth")
	}

	token := ctx.Query("token")
	if token == "" {
		return fiber.NewError(fiber.StatusBadRequest, "token query parameter is required")
	}

	// Пытаемся авторизоваться по токену
	// Сервис вернет данные пользователя или сессию, если токен валиден
	result, err := c.service.VerifyMagicLink(token)
	if err != nil {
		if err == sql.ErrNoRows {
			return fiber.NewError(fiber.StatusUnauthorized, "invalid or expired token")
		}
		// Если токен уже использован или другая ошибка
		return fiber.NewError(fiber.StatusUnauthorized, "verification failed")
	}

	// Устанавливаем куки с сессией (или возвращаем JWT в теле)
	// HttpOnly: true — защита от XSS, Secure: true — только HTTPS
	ctx.Cookie(&fiber.Cookie{
		Name:     "session_token",
		Value:    result.SessionToken,
		Path:     "/",
		MaxAge:   24 * 60 * 60, // 24 часа
		HTTPOnly: true,
		Secure:   ctx.Protocol() == "https",
		SameSite: "lax",
	})

	// Возвращаем данные пользователя (без чувствительных полей)
	// Или делаем редирект: return ctx.Redirect("/dashboard")
	return ctx.JSON(fiber.Map{
		"message": "success",
		"user": fiber.Map{
			"id":       result.UserID,
			"fullname": result.Fullname,
			"email":    result.Email,
			"role":     result.Role,
		},
	})
}
