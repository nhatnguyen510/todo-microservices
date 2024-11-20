package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	kitgrpc "github.com/go-kit/kit/transport/grpc"
	"github.com/go-kit/log"
	pb "github.com/nhatnguyen0510/todo-microservices-app/api/v1/pb/todo"
	"github.com/nhatnguyen0510/todo-microservices-app/pkg/todo"
	"github.com/nhatnguyen0510/todo-microservices-app/pkg/todo/endpoints"
	"github.com/nhatnguyen0510/todo-microservices-app/pkg/todo/transport"
	"github.com/oklog/oklog/pkg/group"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	DefaultHTTPPort = "8081"
	DefaultGRPCPort = "8082"
)

func main() {
	var (
		logger   log.Logger
		httpAddr = net.JoinHostPort("localhost", envString("HTTP_PORT", DefaultHTTPPort))
		grpcAddr = net.JoinHostPort("localhost", envString("GRPC_PORT", DefaultGRPCPort))
	)

	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)

	logger.Log("HTTP", httpAddr)
	logger.Log("GRPC", grpcAddr)

	var (
		service     = todo.NewTodoService()
		endpoints   = endpoints.NewEndpointsSet(service)
		httpHandler = transport.NewHTTPHandler(endpoints, logger)
		grpcServer  = transport.NewGRPCServer(endpoints)
	)

	var g group.Group

	{
		httpListener, err := net.Listen("tcp", httpAddr)
		if err != nil {
			logger.Log("transport", "HTTP", "during", "Listen", "err", err)
			os.Exit(1)
		}

		g.Add(func() error {
			logger.Log("transport", "HTTP", "addr", httpAddr)
			return http.Serve(httpListener, httpHandler)
		}, func(err error) {
			httpListener.Close()
		})
	}

	{
		grpcListener, err := net.Listen("tcp", grpcAddr)
		if err != nil {
			logger.Log("transport", "gRPC", "during", "Listen", "err", err)
			os.Exit(1)
		}

		g.Add(func() error {
			logger.Log("transport", "gRPC", "addr", grpcAddr)
			baseServer := grpc.NewServer(grpc.UnaryInterceptor(kitgrpc.Interceptor))
			pb.RegisterTodoServiceServer(baseServer, grpcServer)
			reflection.Register(baseServer)
			return baseServer.Serve(grpcListener)
		}, func(err error) {
			grpcListener.Close()
		})

	}

	{
		// This function just sits and waits for ctrl-C.
		cancelInterrupt := make(chan struct{})
		g.Add(func() error {
			c := make(chan os.Signal, 1)
			signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
			select {
			case sig := <-c:
				return fmt.Errorf("received signal %s", sig)
			case <-cancelInterrupt:
				return nil
			}
		}, func(error) {
			close(cancelInterrupt)
		})
	}
	logger.Log("exit", g.Run())
}

func envString(env, fallback string) string {
	e := os.Getenv(env)
	if e == "" {
		return fallback
	}
	return e
}
