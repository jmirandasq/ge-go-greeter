package usecase

import (
	"fmt"
	"go-app/internal/domain"
)

type GreeterUseCase struct{}

func NewGreeterUseCase() *GreeterUseCase {
	return &GreeterUseCase{}
}

func (uc *GreeterUseCase) Greet(req domain.GreetRequest) domain.GreetResponse {
	return domain.GreetResponse{
		Message: fmt.Sprintf("¡Hola, %s!", req.Name),
	}
}
