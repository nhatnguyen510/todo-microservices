package endpoints

import (
	"github.com/nhatnguyen0510/todo-microservices-app/internal"
)

type GetTodoRequest struct {
	ID string `json:"id"`
}

type GetTodoResponse struct {
	Err  string        `json:"error,omitempty"`
	Todo internal.Todo `json:"todo"`
}

type CreateTodoRequest struct {
	Title string `json:"title"`
}

type CreateTodoResponse struct {
	Err  string        `json:"error,omitempty"`
	Todo internal.Todo `json:"todo"`
}

type GetTodosRequest struct{}

type GetTodosResponse struct {
	Err   string          `json:"error,omitempty"`
	Todos []internal.Todo `json:"todos"`
}

type UpdateTodoRequest struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

type UpdateTodoResponse struct {
	Err  string        `json:"error,omitempty"`
	Todo internal.Todo `json:"todo"`
}

type DeleteTodoRequest struct {
	ID string `json:"id"`
}

type DeleteTodoResponse struct {
	Err string `json:"error,omitempty"`
}
