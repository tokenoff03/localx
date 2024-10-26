package main

import (
	"context"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"localx"
	"localx/internal/config"
	"localx/internal/handler"
	"localx/internal/repository"
	"localx/internal/services"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Загрузка конфигурации
	cfg, err := config.InitConfig("../config.yml")
	if err != nil {
		logrus.Fatalf("failed to load config: %s", err.Error())
	}
	logrus.SetFormatter(new(logrus.JSONFormatter))
	logrus.Infof("Database URI: %s", cfg.DB.URL)

	// Подключение к базе данных
	db, err := repository.NewPostgresDB(cfg.DB.URL)
	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}
	defer db.Close()

	// Настройки для Supabase

	// Создание сервиса аватаров
	avatarService := services.NewAvatarService("https://kytnrvcpxizrtharguee.supabase.co/storage/v1", "avatars", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJzdXBhYmFzZSIsInJlZiI6Imt5dG5ydmNweGl6cnRoYXJndWVlIiwicm9sZSI6ImFub24iLCJpYXQiOjE3Mjk5MjMzNTMsImV4cCI6MjA0NTQ5OTM1M30.5UGSLi6rjflcFVfPAuWNBjCDWirXAHxEfTClKJXAkHE")

	// Создание репозиториев и сервисов
	repos := repository.NewRepository(db)
	service := services.NewServices(repos)

	// Создание обработчиков
	handler := handler.NewHandler(service, avatarService)

	// Запуск HTTP сервера
	server := new(localx.Server)
	go func() {
		if err := server.Run(cfg.Server.Port, handler.InitRoutes()); err != nil {
			logrus.Fatalf("Error initializing HTTP server: %s", err.Error())
		}
	}()
	logrus.Print("Server Started")

	// Завершение работы
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	logrus.Print("Server Shutting Down")

	if err := server.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occurred on server shutting down: %s", err.Error())
	}
}
