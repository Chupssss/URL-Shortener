package main

import (
	"URL-Shortener/internal/config"
	"URL-Shortener/internal/handler"
	"URL-Shortener/internal/repos"
	"URL-Shortener/internal/routers"
	"URL-Shortener/internal/service"
)

func main() {
	cfg := config.Load()

	repo := repos.NewUrlRepos()
	service := service.NewUrlService(repo)
	handler := handler.NewUrlHandler(service, cfg.BaseURL)

	router := routers.SetupRouter(handler)
	router.Run(":" + cfg.Port)
}
