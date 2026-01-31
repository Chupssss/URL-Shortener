package routers

import (
	"URL-Shortener/internal/handler"

	"github.com/gin-gonic/gin"
)

func SetupRouter(h *handler.UrlHandler) *gin.Engine {
	router := gin.Default()

	// Пинг сервера
	router.GET("/health", handler.Health)

	// Укоротитель
	router.POST("/urls", h.Create)

	// Переадресация
	router.GET("/:shortCode", h.RedirectURL)

	return router
}
