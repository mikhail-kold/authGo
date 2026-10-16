package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	grpcServer "github.com/mikhail-kold/authGo/internal/auth/adapters/grpc"
	"github.com/mikhail-kold/authGo/internal/auth/ports"
	"github.com/mikhail-kold/authGo/internal/auth/service"
	"github.com/mikhail-kold/authGo/internal/common/genproto/auth"
	"github.com/mikhail-kold/authGo/internal/common/logs"
	"github.com/mikhail-kold/authGo/internal/common/server"
	"google.golang.org/grpc"
)

func main() {
	logs.Init()

	app := service.NewApplication()

	go server.RunGRPCServer(func(s *grpc.Server) {
		svc := grpcServer.NewAuthGrpc(app.TokenService)
		auth.RegisterAuthServiceServer(s, svc)
	})

	server.RunHTTPServer(func(router chi.Router) http.Handler {
		handler := ports.HandlerFromMux(ports.NewHttpServer(app), router)

		err := chi.Walk(router, func(method, route string, _ http.Handler, _ ...func(http.Handler) http.Handler) error {
			log.Printf("route: %s %s", method, route)
			return nil
		})
		if err != nil {
			log.Printf("failed to walk routes: %v", err)
		}

		return handler
	})
}
