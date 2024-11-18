package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/vitosotdihaet/map-pinner/pkg/controllers"
	"github.com/vitosotdihaet/map-pinner/pkg/handlers"
	"github.com/vitosotdihaet/map-pinner/pkg/server"
	"github.com/vitosotdihaet/map-pinner/pkg/services"

	_ "github.com/lib/pq"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	logrus.SetLevel(logrus.TraceLevel)

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading .env variables: %s", err.Error())
	}

	postgres, err := controllers.NewPostgresDB(controllers.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USER"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
		Password: os.Getenv("DB_PASSWORD"),
	})

	if err != nil {
		logrus.Fatalf("error while connecting to the database: %s", err.Error())
	}

	database := controllers.NewDatabase(postgres)
	service := services.NewService(database)
	handler := handlers.NewHandler(service)

	server := new(server.Server)

	if err := server.Run(os.Getenv("SERVER_PORT"), handler.InitEndpoints()); err != nil {
		logrus.Fatalf("error while running the server: %s", err.Error())
	}
}
