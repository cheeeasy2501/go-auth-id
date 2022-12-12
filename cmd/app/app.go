package app

import (
	"context"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"

	cfg "github.com/cheeeasy2501/auth-id/config"
	ctlr "github.com/cheeeasy2501/auth-id/internal/transport/http/v1/controller"
	srvs "github.com/cheeeasy2501/auth-id/internal/service"

	"github.com/cheeeasy2501/auth-id/pkg/server"
)

func Run(ctx context.Context, log *log.Logger, config *cfg.Config, conn *gorm.DB) {

	httpServer := server.NewHTTPServer()
	router := httpServer.GetRouter()
	services := srvs.NewService(config, conn)
	controllers := ctlr.NewController(services)

	v1 := router.Group("/v1")
	{
		controllers.RegisterRoutes(v1)
	}

	httpServer.StartHTTPServer()
}
