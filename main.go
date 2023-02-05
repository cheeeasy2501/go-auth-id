package main

import (
	//"os"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
   ctx "context"
   "os/signal"
	"syscall"

	"github.com/cheeeasy2501/auth-id/cmd/app"
   cfg "github.com/cheeeasy2501/auth-id/config"
   "github.com/cheeeasy2501/auth-id/pkg/database"
)

func main() {
	// инициализируем логгер
	logger := log.New()

	// инициализируем .env
	if err := godotenv.Load("env.local"); err != nil {
		logger.Fatal("No .env file found")
	}

   logger.Infoln(".env variables is loaded")

   // инициализируем базу данных и конфиг
   config:= cfg.NewConfig()
   db := database.NewDB(config.Database)

   conn, err := db.OpenConnection()
   if err != nil {
      logger.Fatal("Connection isn't opened!", err)
   }
   defer db.CloseConnection()
   logger.Infoln("Database connection is opened")

   ctx, cancel := signal.NotifyContext(ctx.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	//  стартуем приложение
   logger.Infoln("Starting application")
	app.Run(ctx, logger, config, conn)
   logger.Infoln("Application is started")

   <-ctx.Done()
}
