package endpoints

import (
	"context"
	"os"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/log"
	"github.com/nhatnguyen0510/todo-microservices-app/pkg/todo"
)

type Set struct {
	CreateTodoEndpoint endpoint.Endpoint
	GetTodoEndpoint    endpoint.Endpoint
	GetTodosEndpoint   endpoint.Endpoint
	UpdateTodoEndpoint endpoint.Endpoint
	DeleteTodoEndpoint endpoint.Endpoint
}

func NewEndpointsSet(s todo.Service) Set {
	return Set{
		CreateTodoEndpoint: MakeCreateTodoEndpoint(s),
		GetTodoEndpoint:    MakeGetTodoEndpoint(s),
		GetTodosEndpoint:   MakeGetTodosEndpoint(s),
		UpdateTodoEndpoint: MakeUpdateTodoEndpoint(s),
		DeleteTodoEndpoint: MakeDeleteTodoEndpoint(s),
	}
}

func MakeDeleteTodoEndpoint(s todo.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteTodoRequest)
		err := s.DeleteTodo(ctx, req.ID)

		if err != nil {
			return DeleteTodoResponse{Err: err.Error()}, nil

		}
		return DeleteTodoResponse{Err: ""}, nil
	}
}

func MakeUpdateTodoEndpoint(s todo.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateTodoRequest)
		todo, err := s.UpdateTodo(ctx, req.ID, req.Title)

		if err != nil {
			return UpdateTodoResponse{Err: err.Error()}, nil

		}

		return UpdateTodoResponse{Err: "", Todo: todo}, nil
	}
}

func MakeGetTodosEndpoint(s todo.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		todos, err := s.GetTodos(ctx)

		if err != nil {
			return GetTodosResponse{Err: err.Error()}, nil
		}
		return GetTodosResponse{Err: "", Todos: todos}, nil
	}
}

func MakeGetTodoEndpoint(s todo.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetTodoRequest)
		todo, err := s.GetTodo(ctx, req.ID)

		if err != nil {
			return GetTodoResponse{Err: err.Error()}, nil
		}
		return GetTodoResponse{Err: "", Todo: todo}, nil
	}
}

func MakeCreateTodoEndpoint(s todo.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateTodoRequest)
		todo, err := s.CreateTodo(ctx, req.Title)

		if err != nil {
			return CreateTodoResponse{Err: err.Error()}, nil
		}
		return CreateTodoResponse{Err: "", Todo: todo}, nil
	}
}

var logger log.Logger

func init() {
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
}
