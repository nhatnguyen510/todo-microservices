package todo

import (
	"context"
	"os"

	"github.com/nhatnguyen0510/todo-microservices-app/internal"

	"github.com/go-kit/log"
)

type Service interface {
	CreateTodo(ctx context.Context, title string) (internal.Todo, error)
	GetTodo(ctx context.Context, id string) (internal.Todo, error)
	GetTodos(ctx context.Context) ([]internal.Todo, error)
	UpdateTodo(ctx context.Context, id string, title string) (internal.Todo, error)
	DeleteTodo(ctx context.Context, id string) error
}

type todoService struct{}

func NewTodoService() Service {
	return &todoService{}
}

func (s *todoService) CreateTodo(ctx context.Context, title string) (internal.Todo, error) {
	return internal.Todo{}, nil
}

func (s *todoService) GetTodo(ctx context.Context, id string) (internal.Todo, error) {
	return internal.Todo{}, nil
}

func (s *todoService) GetTodos(ctx context.Context) ([]internal.Todo, error) {
	return []internal.Todo{}, nil
}

func (s *todoService) UpdateTodo(ctx context.Context, id string, title string) (internal.Todo, error) {
	return internal.Todo{}, nil
}

func (s *todoService) DeleteTodo(ctx context.Context, id string) error {
	return nil
}

var logger log.Logger

func init() {
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
}
