package app

import (
	"context"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"

	cfg "github.com/cheeeasy2501/auth-id/config"
	srvs "github.com/cheeeasy2501/auth-id/internal/service"

	// mwr "github.com/cheeeasy2501/auth-id/internal/transport/http/middleware"
	ctlr "github.com/cheeeasy2501/auth-id/internal/transport/http/v1/controller"

	"github.com/cheeeasy2501/auth-id/pkg/server"
)

func Run(ctx context.Context, log *log.Logger, config *cfg.Config, conn *gorm.DB) {

	httpServer := server.NewHTTPServer()
	router := httpServer.GetRouter()
	services := srvs.NewService(config, conn)
	controllers := ctlr.NewController(services)
	// middlewares := mwr.NewMiddleware(services.Authorization)
	v1 := router.Group("/v1")
	{
		// controllers.RegisterRoutes(routes)

		// функционал
		// востановление доступа к учетной записи (пароль)
		// изменение данных ползователя
		// методы:
		// получение access jwt
		// получение refresh jwt
		// обновление access jwt
		// обновление refresh jwt
		// реализация авторизации по токену
		//  1. отправляем access token
		//  2. если access token невалидный - возвращаем на фронт ошибку
		//  3. берем refresh токен, отсылаем на ендпоинт refresh и получаем новый access
		//  4. отсылаем повторно access
		// реализация по доступу к email-sender
		//  1. посылаем напрямую в email-sender
		//  2. email-sender отсылает запрос в auth-id для и проверяет токен
		//  3. ????
		//  4. Profit
		// v1.POST("/jwt", middlewares.Authorize)

	}

	auth := v1.Group("/auth")
	{
		auth.POST("/login", controllers.Authorization.LoginByEmail)        //выдать access и refresh(?)
		auth.POST("/registration", controllers.Authorization.Registration) // отправляю письмо на email?
		auth.POST("/refresh-tokens", controllers.Authorization.RefreshTokens)
	}

	httpServer.StartHTTPServer()
}
