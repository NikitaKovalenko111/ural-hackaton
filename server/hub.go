package hubController

import "github.com/gofiber/fiber/v2"

type HTTP struct {
	app                     *fiber.App
	debtorController        *hubController.DebtorController
	adminController         *userController.AdminController
	deptorMessageController *requestController.DebtorMesController
	authMiddleware          func(c *fiber.Ctx) error
}
