package transport

import (
	"context"

	grpctransport "github.com/go-kit/kit/transport/grpc"
	todo "github.com/nhatnguyen0510/todo-microservices-app/api/v1/pb/todo"
	"github.com/nhatnguyen0510/todo-microservices-app/pkg/todo/endpoints"
)

type TodoHandler string

const (
	CreateTodoHandler TodoHandler = "CreateTodo"
	GetTodoHandler    TodoHandler = "GetTodo"
	GetTodosHandler   TodoHandler = "GetTodos"
	UpdateTodoHandler TodoHandler = "UpdateTodo"
	DeleteTodoHandler TodoHandler = "DeleteTodo"
)

type grpcServer struct {
	todo.UnimplementedTodoServiceServer
	handlers map[TodoHandler]grpctransport.Handler
}

func NewGRPCServer(ep endpoints.Set) todo.TodoServiceServer {
	server := &grpcServer{
		handlers: map[TodoHandler]grpctransport.Handler{
			CreateTodoHandler: grpctransport.NewServer(
				ep.CreateTodoEndpoint,
				decodeCreateTodoRequest,
				encodeCreateTodoResponse,
			),
			GetTodoHandler: grpctransport.NewServer(
				ep.GetTodoEndpoint,
				decodeGetTodoRequest,
				encodeGetTodoResponse,
			),
			GetTodosHandler: grpctransport.NewServer(
				ep.GetTodosEndpoint,
				decodeGetTodosRequest,
				encodeGetTodosResponse,
			),
			UpdateTodoHandler: grpctransport.NewServer(
				ep.UpdateTodoEndpoint,
				decodeUpdateTodoRequest,
				encodeUpdateTodoResponse,
			),
			DeleteTodoHandler: grpctransport.NewServer(
				ep.DeleteTodoEndpoint,
				decodeDeleteTodoRequest,
				encodeDeleteTodoResponse,
			),
		},
	}

	return server
}

func (s *grpcServer) handleRequest(ctx context.Context, req interface{}, handlerName TodoHandler) (interface{}, error) {
	_, resp, err := s.handlers[handlerName].ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *grpcServer) CreateTodo(ctx context.Context, req *todo.CreateTodoRequest) (*todo.CreateTodoResponse, error) {
	resp, err := s.handleRequest(ctx, req, CreateTodoHandler)
	if err != nil {
		return nil, err
	}
	return resp.(*todo.CreateTodoResponse), nil
}

func (s *grpcServer) GetTodo(ctx context.Context, req *todo.GetTodoRequest) (*todo.GetTodoResponse, error) {
	resp, err := s.handleRequest(ctx, req, GetTodoHandler)
	if err != nil {
		return nil, err
	}
	return resp.(*todo.GetTodoResponse), nil
}

func (s *grpcServer) GetTodos(ctx context.Context, req *todo.GetTodosRequest) (*todo.GetTodosResponse, error) {
	resp, err := s.handleRequest(ctx, req, GetTodosHandler)
	if err != nil {
		return nil, err
	}
	return resp.(*todo.GetTodosResponse), nil
}

func (s *grpcServer) UpdateTodo(ctx context.Context, req *todo.UpdateTodoRequest) (*todo.UpdateTodoResponse, error) {
	resp, err := s.handleRequest(ctx, req, UpdateTodoHandler)
	if err != nil {
		return nil, err
	}
	return resp.(*todo.UpdateTodoResponse), nil
}

func (s *grpcServer) DeleteTodo(ctx context.Context, req *todo.DeleteTodoRequest) (*todo.DeleteTodoResponse, error) {
	resp, err := s.handleRequest(ctx, req, DeleteTodoHandler)
	if err != nil {
		return nil, err
	}
	return resp.(*todo.DeleteTodoResponse), nil
}

func decodeCreateTodoRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*todo.CreateTodoRequest)
	return endpoints.CreateTodoRequest{Title: req.Title}, nil
}

func encodeCreateTodoResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoints.CreateTodoResponse)
	return &todo.CreateTodoResponse{Todo: &todo.Todo{Id: resp.Todo.ID, Title: resp.Todo.Title}}, nil
}

func decodeGetTodoRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*todo.GetTodoRequest)
	return endpoints.GetTodoRequest{ID: req.Id}, nil
}

func encodeGetTodoResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoints.GetTodoResponse)
	return &todo.GetTodoResponse{Todo: &todo.Todo{Id: resp.Todo.ID, Title: resp.Todo.Title}}, nil
}

func decodeGetTodosRequest(_ context.Context, _ interface{}) (interface{}, error) {
	return endpoints.GetTodosRequest{}, nil
}

func encodeGetTodosResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoints.GetTodosResponse)
	todos := make([]*todo.Todo, len(resp.Todos))
	for i, t := range resp.Todos {
		todos[i] = &todo.Todo{Id: t.ID, Title: t.Title}
	}
	return &todo.GetTodosResponse{Todos: todos}, nil
}

func decodeUpdateTodoRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*todo.UpdateTodoRequest)
	return endpoints.UpdateTodoRequest{ID: req.Id, Title: req.Title}, nil
}

func encodeUpdateTodoResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoints.UpdateTodoResponse)
	return &todo.UpdateTodoResponse{Todo: &todo.Todo{Id: resp.Todo.ID, Title: resp.Todo.Title}}, nil
}

func decodeDeleteTodoRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*todo.DeleteTodoRequest)
	return endpoints.DeleteTodoRequest{ID: req.Id}, nil
}

func encodeDeleteTodoResponse(_ context.Context, _ interface{}) (interface{}, error) {
	return &todo.DeleteTodoResponse{}, nil
}
