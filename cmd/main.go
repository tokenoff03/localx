package main

import (
	"context"
	"localx"
	"localx/internal/handler"
	"localx/internal/repository"
	"localx/internal/services"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(viper.GetString("db.url"))
	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	service := services.NewServices(repos)
	handler := handler.NewHandler(service)

	server := new(localx.Server)
	go func() {
		if err := server.Run(viper.GetString("port"), handler.InitRoutes()); err != nil {
			logrus.Fatalf("Error to initializing http server: %s", err.Error())
		}
	}()

	logrus.Print("HotelGO Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("HotelGO Shutting Down")
	if err := server.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		logrus.Errorf("Error occured on db connection close: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("../internal/configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
