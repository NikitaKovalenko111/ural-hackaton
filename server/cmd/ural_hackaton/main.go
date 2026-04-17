package main

import (
	"log"
	"ural-hackaton/internal/app"
	"ural-hackaton/internal/config"
)

func main() {
	cfg := config.MustLoad()

	if err := app.Run(cfg); err != nil {
		log.Fatal(err)
	}
}
