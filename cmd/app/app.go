package app

import (
	"context"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"

	cfg "github.com/cheeeasy2501/auth-id/config"
	srvs "github.com/cheeeasy2501/auth-id/internal/service"

	gen "github.com/cheeeasy2501/auth-id/gen"
	auth_grpc "github.com/cheeeasy2501/auth-id/internal/transport/grpc/v1/authorization"
	ctlr "github.com/cheeeasy2501/auth-id/internal/transport/http/v1/controller"
	mwr "github.com/cheeeasy2501/auth-id/internal/transport/http/v1/middleware"

	grpcserver "github.com/cheeeasy2501/auth-id/pkg/grpc-server"
	"github.com/cheeeasy2501/auth-id/pkg/server"
)

// Запуск приложения
func Run(ctx context.Context, l *log.Logger, cfg *cfg.Config, conn *gorm.DB) {
	services := srvs.NewService(cfg, conn)

	go func() {
		grpcSrv := grpcserver.NewGRPCServer(cfg.GRPC)
		srv := auth_grpc.NewAuthorizationGRPCServer(services)
		gen.RegisterAuthorizationServiceServer(grpcSrv.GetInstance(), srv)

		err := grpcSrv.Run()
		if err != nil {
			panic("GRPC isn't started!")
		}

		l.Infoln("GRPC started on port " + grpcSrv.GetConfig().GetPort())
	}()

	go func() {
		httpServer := server.NewHTTPServer(cfg.HTTP)
		router := httpServer.GetRouter()

		controllers := ctlr.NewController(services)
		middleware := mwr.NewMiddleware(services.Authorization)

		auth := router.Group("/v1/auth")
		{
			auth.POST("/login", controllers.Authorization.LoginByEmail)
			auth.POST("/registration", controllers.Authorization.Registration) // отправляю письмо на email?
			auth.POST("/refresh-token", middleware.Jwtm.CheckRefreshToken(), controllers.Authorization.RefreshToken)
		}

		err := httpServer.Run()
		if err != nil {
			panic("HTTP isn't started!")
		}

		l.Infoln("HTTP started on port " + httpServer.GetConfig().GetPort())
	}()
}
