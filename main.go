package main

import (
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"

	"github.com/cheeeasy2501/cmd/app"
)

func main() {
	// инициализируем логгер
	logger := log.New()
	// инициализируем .env
	if err := godotenv.Load(); err != nil {

	}
	// Инициализируем  репозитории, сервисы, мб GRPC-сервис

	//  стартуем приложение
	app.Run(logger)
}
