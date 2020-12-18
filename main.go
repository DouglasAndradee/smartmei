package main

import (
	"context"
	"os"
	"os/signal"
	"time"

	"github.com/DouglasAndradee/smartmei/controller"
	"github.com/DouglasAndradee/smartmei/database"
	"github.com/DouglasAndradee/smartmei/repository"
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

	r := &repository.Repository{}
	r.Session = database.Session()

	api.GET("/user/:id", controller.GetUser(r))
	api.POST("/user", controller.InsertUser(r))
	api.POST("/book", controller.InsertBookToUser(r))

	api.PUT("/book/lend", controller.LendBook(r))
	api.PUT("/book/return", controller.ReturnBook(r))

	go func() {
		if err := api.Start(":5000"); err != nil {
			api.Logger.Warn("DESLIGANDO O SERVIÇO")
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := api.Shutdown(ctx); err != nil {
		api.Logger.Fatal(err)
	}
}
