package main

import (
	"context"
	"log"

	"go-app/internal/infraestructure/entrypoints"
	"go-app/internal/infraestructure/telemetry"
	"go-app/internal/usecase"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

func main() {
	ctx := context.Background()

	tp, err := telemetry.InitTracer(ctx)
	if err != nil {
		log.Fatalf("failed to initialize tracer: %v", err)
	}
	defer tp.Shutdown(ctx)

	uc := usecase.NewGreeterUseCase()
	handler := entrypoints.NewGreeterHandler(uc)

	router := gin.Default()
	router.Use(otelgin.Middleware("ge-go-greeter"))

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "API is running."})
	})

	router.POST("/greet", handler.Greet)

	router.GET("/test-401", func(c *gin.Context) {
		c.JSON(401, gin.H{"error": "Unauthorized"})
	})

	router.GET("/test-error", func(c *gin.Context) {
		log.Println("ERROR: fallo critico en la aplicacion.")
		c.JSON(500, gin.H{"error": "Internal Server Error"})
	})

	router.Run()
}
