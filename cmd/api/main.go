package main

import (
	"go-app/internal/infraestructure/entrypoints"
	"go-app/internal/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	uc := usecase.NewGreeterUseCase()
	handler := entrypoints.NewGreeterHandler(uc)

	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "API is running."})
	})

	router.POST("/greet", handler.Greet)

	router.Run()
}
