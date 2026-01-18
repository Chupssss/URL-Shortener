package main

import (
	"URL-Shortener/internal/config"
	"URL-Shortener/internal/routers"
)

func main() {
	cfg := config.Load()

	router := routers.SetupRouter()
	router.Run(":" + cfg.Port)
}
