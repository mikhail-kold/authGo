package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/mmMikhailS/DeepBoard/internal/auth/service"
	"github.com/mmMikhailS/DeepBoard/internal/common/logs"
	"github.com/mmMikhailS/DeepBoard/internal/common/server"
)

func main() {
	logs.Init()

	app := service.NewApplication()

	// TODO:
	//go server.RunGRPCServer(func(s *grpc.Server) {
	//	svc := grpcServer.NewAuthGrpc(app.TokenService)
	//	auth.RegisterAuthServiceServer(s, svc)
	//})

	server.RunHTTPServer(func(router chi.Router) http.Handler {
		return ports.HandlerFromMux(ports.NewHttpServer(app), router)
	})
}
