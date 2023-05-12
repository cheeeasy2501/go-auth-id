package app

import (
	"context"
	"net"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"

	cfg "github.com/cheeeasy2501/auth-id/config"
	srvs "github.com/cheeeasy2501/auth-id/internal/service"

	auth_gen "github.com/cheeeasy2501/auth-id/gen/authorization"
	auth_grpc "github.com/cheeeasy2501/auth-id/internal/transport/grpc/v1/authorization"
	ctlr "github.com/cheeeasy2501/auth-id/internal/transport/http/v1/controller"
	mwr "github.com/cheeeasy2501/auth-id/internal/transport/http/v1/middleware"

	"github.com/cheeeasy2501/auth-id/pkg/server"
	"google.golang.org/grpc"
)

// Запуск приложения
func Run(ctx context.Context, l *log.Logger, config *cfg.Config, conn *gorm.DB) {
	httpServer := server.NewHTTPServer(config.HTTP)
	router := httpServer.GetRouter()
	services := srvs.NewService(config, conn)
	controllers := ctlr.NewController(services)
	middleware := mwr.NewMiddleware(services.Authorization)

	auth := router.Group("/v1/auth")
	{
		auth.POST("/login", controllers.Authorization.LoginByEmail)
		auth.POST("/registration", controllers.Authorization.Registration) // отправляю письмо на email?
		auth.POST("/refresh-token", middleware.Jwtm.CheckRefreshToken(), controllers.Authorization.RefreshToken)
	}

	go func() {
		err := startGRPCServer()
		if err != nil {
			panic("GRPC isn't started!")
		}

		l.Infoln("GRPC started on port 1000")
	}()
	go func() {
		err := httpServer.StartHTTPServer()
		if err != nil {
			panic("HTTP isn't started!")
		}

		l.Infoln("GRPC started on port " + httpServer.GetConfig().GetPort())
	}()
}

// Запуск GRPC
func startGRPCServer() error {
	grpcServer := grpc.NewServer()
	srv := &auth_grpc.AuthorizationGRPCServer{}
	auth_gen.RegisterAuthorizationServiceServer(grpcServer, srv)

	l, err := net.Listen("tcp", ":1000")
	if err != nil {
		return err
	}

	if err := grpcServer.Serve(l); err != nil {
		return err
	}

	return nil
}
