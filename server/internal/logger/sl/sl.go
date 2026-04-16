package sl

import (
	"log/slog"
	"os"
	"ural-hackaton/server/internal/types"

	"github.com/lmittmann/tint"
)

func InitLogger(env string) *slog.Logger {
	var slogLevel slog.Level
	switch env {
	case types.EnvLocal:
		slogLevel = slog.LevelDebug
	case types.EnvDev:
		slogLevel = slog.LevelDebug
	case types.EnvProd:
		slogLevel = slog.LevelInfo
	}
	logger := slog.New(tint.NewHandler(os.Stdout, &tint.Options{
		Level:     slogLevel,
		AddSource: true,
	}))
	return logger
}
