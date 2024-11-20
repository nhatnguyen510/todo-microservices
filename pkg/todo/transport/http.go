package transport

import (
	"context"
	"encoding/json"
	"net/http"
	"os"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/go-kit/log"
	"github.com/nhatnguyen0510/todo-microservices-app/internal/util"
	"github.com/nhatnguyen0510/todo-microservices-app/pkg/todo/endpoints"
)

func NewHTTPHandler(ep endpoints.Set, logger log.Logger) http.Handler {
	m := http.NewServeMux()
	m.Handle("/todo/create", httptransport.NewServer(
		ep.CreateTodoEndpoint,
		decodeJSONRequest(endpoints.CreateTodoRequest{}),
		encodeResponse,
	))
	m.Handle("/todos", httptransport.NewServer(
		ep.GetTodosEndpoint,
		decodeJSONRequest(endpoints.GetTodosRequest{}),
		encodeResponse,
	))
	m.Handle("/todo/get/", httptransport.NewServer(
		ep.GetTodoEndpoint,
		decodeJSONRequest(endpoints.GetTodoRequest{}),
		encodeResponse,
	))
	m.Handle("/todo/update/", httptransport.NewServer(
		ep.UpdateTodoEndpoint,
		decodeJSONRequest(endpoints.UpdateTodoRequest{}),
		encodeResponse,
	))
	m.Handle("/todo/delete/", httptransport.NewServer(
		ep.DeleteTodoEndpoint,
		decodeJSONRequest(endpoints.DeleteTodoRequest{}),
		encodeResponse,
	))

	return m
}

func decodeJSONRequest(req interface{}) httptransport.DecodeRequestFunc {
	return func(_ context.Context, r *http.Request) (interface{}, error) {
		if r.ContentLength == 0 {
			logger.Log("Request with no body")
			return req, nil
		}
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			return nil, err
		}
		return req, nil
	}
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(error); ok && e != nil {
		encodeError(ctx, e, w)
		return nil
	}
	return json.NewEncoder(w).Encode(response)
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	switch err {
	case util.ErrUnknown:
		w.WriteHeader(http.StatusNotFound)
	case util.ErrInvalidArgument:
		w.WriteHeader(http.StatusBadRequest)
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

var logger log.Logger

func init() {
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
}
