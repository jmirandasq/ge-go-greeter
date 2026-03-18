package domain

type GreeterPort interface {
	Greet(req GreetRequest) GreetResponse
}
