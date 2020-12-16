package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/DouglasAndradee/smartmei/database"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	api := echo.New()
	api.HideBanner = true

	api.Use(middleware.CORS())
	api.Use(middleware.Secure())
	api.Use(middleware.Recover())
	api.Use(middleware.RequestID())

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database.Session()

	go func() {
		if err := api.Start(":5000"); err != nil {
			api.Logger.Info("shutting down the server")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := api.Shutdown(ctx); err != nil {
		api.Logger.Fatal(err)
	}

}
