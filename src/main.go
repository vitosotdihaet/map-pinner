package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/vitosotdihaet/map-pinner/package/controllers"
	"github.com/vitosotdihaet/map-pinner/package/handlers"
	"github.com/vitosotdihaet/map-pinner/package/server"
	"github.com/vitosotdihaet/map-pinner/package/services"

	_ "github.com/lib/pq"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("error while getting the config: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading .env variables: %s", err.Error())
	}


	postgres, err := controllers.NewPostgresDB(controllers.Config{
		Host: viper.GetString("db.host"),
		Port: viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName: viper.GetString("db.dbname"),
		SSLMode: viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})

	if err != nil {
		logrus.Fatalf("error while connecting to the database: %s", err.Error())
	}


	database := controllers.NewDatabase(postgres)
	service := services.NewService(database)
	handler := handlers.NewHandler(service)

	server := new(server.Server)

	if err := server.Run(viper.GetString("port"), handler.InitEndpoints()); err != nil {
		logrus.Fatalf("error while running the server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
