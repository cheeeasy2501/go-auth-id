package app

import (
	"context"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"

	cfg "github.com/cheeeasy2501/auth-id/config/app"
	ctlr "github.com/cheeeasy2501/auth-id/internal/controller/app"
	srvs "github.com/cheeeasy2501/auth-id/internal/service/app"

	"github.com/cheeeasy2501/auth-id/package/server"
)

func Run(ctx context.Context, log *log.Logger, config *cfg.Config, conn *gorm.DB) {

	httpServer := server.NewHTTPServer()
	router := httpServer.GetRouter()
	services := srvs.NewService(config, conn)
	controllers := ctlr.NewController(services)

	v1 := router.Group("/v1")
	{
		v1.POST("/login", controllers.Authorization.LoginByEmail)
		// v1.POST("/register", controllers.Authorization.Register)
	}

	httpServer.StartHTTPServer()
}
