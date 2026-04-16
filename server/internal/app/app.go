package app

import (
	"ural-hackaton/internal/config"
	"ural-hackaton/internal/logger/sl"
)

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.
func Run(cfg *config.Config) {
	logger := sl.InitLogger(cfg.Env)

	logger.Info("Logger is enabled")
	logger.Debug("Debug is enabled")

}
