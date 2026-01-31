package routers

import (
	"URL-Shortener/internal/handler"
	"URL-Shortener/internal/repos"
	"URL-Shortener/internal/service"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Пинг сервера
	router.GET("/health", handler.Health)

	repo := repos.NewUrlRepos()
	service := service.NewUrlService(repo)
	handler := handler.NewUrlHandler(service)

	// Укоротитель
	router.POST("/urls", handler.Create)

	// Переадресация
	router.GET("/:shortCode", handler.RedirectURL)

	return router
}
