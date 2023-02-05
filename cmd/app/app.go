package app

import (
	"context"
	"net"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"

	cfg "github.com/cheeeasy2501/auth-id/config"
	srvs "github.com/cheeeasy2501/auth-id/internal/service"

	ctlr "github.com/cheeeasy2501/auth-id/internal/transport/http/v1/controller"
	mwr "github.com/cheeeasy2501/auth-id/internal/transport/http/v1/middleware"

	pb "github.com/cheeeasy2501/auth-id/pb/authorization"
	"github.com/cheeeasy2501/auth-id/pkg/server"
	"google.golang.org/grpc"
)

func Run(ctx context.Context, log *log.Logger, config *cfg.Config, conn *gorm.DB) {

	httpServer := server.NewHTTPServer()
	router := httpServer.GetRouter()
	services := srvs.NewService(config, conn)
	controllers := ctlr.NewController(services)
	middleware := mwr.NewMiddleware(services.Authorization)

	auth := router.Group("/v1/auth")
	{
		auth.POST("/login", controllers.Authorization.LoginByEmail)
		auth.POST("/registration", controllers.Authorization.Registration) // отправляю письмо на email?
		auth.POST("/refresh-tokens", middleware.Jwtm.Authorize(), controllers.Authorization.RefreshTokens)
	}

	go func() {
		runGRPCServer()
	}()

	httpServer.StartHTTPServer()
}

func runGRPCServer() {
	grpcServer := grpc.NewServer()
	srv := &pb.AuthorizeGRPCServer{}
	pb.RegisterAuthServiceServer(grpcServer, srv)

	l, err := net.Listen("tcp", ":9091")
	if err != nil {
		log.Fatal(err)
	}

	if err := grpcServer.Serve(l); err != nil {
		log.Fatal(err)
	}
}
