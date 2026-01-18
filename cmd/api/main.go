package main

import (
	"URL-Shortener/internal/routers"
)

func main() {
	router := routers.SetupRouter()
	router.Run(":8080")
}
