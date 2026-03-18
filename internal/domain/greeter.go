package domain

type GreetRequest struct {
	Name string `json:"name" binding:"required"`
}

type GreetResponse struct {
	Message string `json:"message"`
}
