package main

import (
	"context"
	"fmt"
	"market_place"
	"market_place/config"
	"market_place/models"
	"market_place/pkg/handler"
	"market_place/pkg/repository"
	"market_place/pkg/service"
	"os"
	"os/signal"
	"syscall"

	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	config.DB.AutoMigrate(&models.User{}, &models.Category{}, &models.Product{})
	db := config.ConnectDB()
	defer config.Disconnect(db)

	repository := repository.NewRepository(db)
	service := service.NewService(repository)
	handlers := handler.NewHandler(service)
	srv := new(market_place.Server)

	go func() {
		if err := srv.Run(os.Getenv("PORT"), handlers.InitRoutes()); err != nil {
			logrus.Fatalf("error server connection: %s", err.Error())
		}
	}()
	fmt.Println("App started...")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	fmt.Println("Server shutting down...")
	if err := srv.Shutdown(context.Background()); err != nil {
		fmt.Errorf("error occured on server shutting down: %s", err.Error())
	}

}
