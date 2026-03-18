package entrypoints

import (
	"go-app/internal/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GreeterHandler struct {
	useCase domain.GreeterPort
}

func NewGreeterHandler(useCase domain.GreeterPort) *GreeterHandler {
	return &GreeterHandler{useCase: useCase}
}

func (h *GreeterHandler) Greet(c *gin.Context) {
	var req domain.GreetRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response := h.useCase.Greet(req)
	c.JSON(http.StatusOK, response)
}
