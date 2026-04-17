package app

import (
	"fmt"
	"ural-hackaton/internal/config"
	"ural-hackaton/internal/logger/sl"
	"ural-hackaton/internal/middleware"
	"ural-hackaton/internal/services"
	"ural-hackaton/internal/storage"
	"ural-hackaton/internal/storage/repositories"
	"ural-hackaton/internal/transport/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.
func Run(cfg *config.Config) error {
	logger := sl.InitLogger(cfg.Env)

	logger.Info("Logger is enabled")
	logger.Debug("Debug is enabled")

	db := storage.Connect(cfg)

	logger.Info("Successfully connected to database!")

	storage := storage.Init(db)

	logger.Info("Successfully inited storage!")

	storage.Prepare()

	logger.Info("Successfully prepared db!")

	repos := repositories.InitRepositories(storage)

	logger.Info("Successfully inited repositories!")

	services := services.Init(repos, cfg)

	logger.Info("Successfully inited services!")

	app := fiber.New(fiber.Config{
		StrictRouting: true,
		WriteTimeout:  cfg.HTTPServer.Timeout,
		IdleTimeout:   cfg.HTTPServer.IdleTimeout,
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:5173,http://127.0.0.1:5173,http://localhost:5174,http://127.0.0.1:5174,http://localhost:3000,http://127.0.0.1:3000",
		AllowMethods:     "GET,POST,PUT,PATCH,DELETE,OPTIONS",
		AllowHeaders:     "Origin,Content-Type,Accept,Authorization",
		ExposeHeaders:    "Content-Length,Content-Type",
		AllowCredentials: true,
	}))

	app.Use(middleware.NewLogger(logger))

	http := http.Init(services, app)

	http.Start()

	logger.Info("Starting server on " + cfg.HTTPServer.Address)
	if err := app.Listen(cfg.HTTPServer.Address); err != nil {
		return fmt.Errorf("fiber listen failed on %s: %w", cfg.HTTPServer.Address, err)
	}

	return nil
}
