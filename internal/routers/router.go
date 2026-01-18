package routers

import (
	"URL-Shortener/internal/handler"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	//Эндпоинт "Работа сервера"
	router.GET("/health", handler.Health)

	return router
}
