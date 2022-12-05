package main

import (
	//"os"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"

	app "github.com/cheeeasy2501/auth-id/cmd/app"
   cfg "github.com/cheeeasy2501/auth-id/config/app"
   "github.com/cheeeasy2501/auth-id/package/database"
)

func main() {
	// инициализируем логгер
	logger := log.New()

	// инициализируем .env
	if err := godotenv.Load(); err != nil {
		logger.Fatal("No .env file found")
	}

   // инициализируем базу данных и конфиг
   config:= cfg.NewConfig()
   db := database.NewDB(config.Database)

   _, err := db.OpenConnection()
   if err != nil {
      logger.Fatal("Connection isn't opened!")
   }
   defer db.CloseConnection()
	// Инициализируем  репозитории, сервисы, мб GRPC-сервис
  // services := NewServices()
	//  стартуем приложение
	app.Run(logger, config)
}

func NewServices() {
	panic("unimplemented")
}
