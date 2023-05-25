package app

import (
	"context"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"

	cfg "github.com/cheeeasy2501/auth-id/config"
	srvs "github.com/cheeeasy2501/auth-id/internal/service"

	ctlr "github.com/cheeeasy2501/auth-id/internal/transport/http/v1/controller"
	mwr "github.com/cheeeasy2501/auth-id/internal/transport/http/v1/middleware"

	auth_grpc "github.com/cheeeasy2501/auth-id/internal/transport/grpc/v1/authorization"
	"github.com/cheeeasy2501/auth-id/pkg/server"
)

// Запуск приложения
func Run(ctx context.Context, l *log.Logger, config *cfg.Config, conn *gorm.DB) {
	services := srvs.NewService(config, conn)
	controllers := ctlr.NewController(services)
	middleware := mwr.NewMiddleware(services.Authorization)

	grpcServer, err := auth_grpc.NewAuthorizationGRPCServer(config.GRPC, services)
	if err != nil {
		return
	}

	err = grpcServer.Run(services)
	if err != nil {
		panic("GRPC isn't started!")
	}

	l.Infoln("GRPC started on address " + config.GRPC.GetAddr())

	httpServer := server.NewHTTPServer(config.HTTP)
	router := httpServer.GetRouter()
	v1 := router.Group("/v1")
	{
		publicRoute := v1.Group("/auth")
		{
			publicRoute.POST("/login", controllers.Authorization.LoginByEmail)
			publicRoute.POST("/registration", controllers.Authorization.Registration) // отправляю письмо на email?
			publicRoute.POST("/refresh-token", middleware.Jwtm.CheckRefreshToken(), controllers.Authorization.RefreshToken)
		}

		protectedRoute := v1.Group("/mail").Use(middleware.Jwtm.Authorize())
		{
			protectedRoute.POST("/send", )
		} 
		
	}


	err = httpServer.Run()
	if err != nil {
		panic("HTTP isn't started!")
	}

	l.Infoln("HTTP started on address " + config.HTTP.GetAddr())

	l.Infoln("Servers has been started...")
}
