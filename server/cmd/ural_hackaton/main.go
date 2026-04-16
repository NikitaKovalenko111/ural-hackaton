package main

import (
	"ural-hackaton/internal/app"
	"ural-hackaton/internal/config"
)

func main() {
	cfg := config.MustLoad()

	app.Run(cfg)
}
