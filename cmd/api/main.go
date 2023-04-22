package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/vuho-pg/base/cmd/api/controller"
	"github.com/vuho-pg/base/pkg/config"
	"github.com/vuho-pg/base/pkg/connection"
	"github.com/vuho-pg/base/pkg/custom_middleware"
	"github.com/vuho-pg/base/pkg/log"
	"github.com/vuho-pg/base/pkg/repository"
	"github.com/vuho-pg/base/pkg/service"
	"go.uber.org/dig"
)

func main() {
	container := dig.New()
	if err := config.Provide(container); err != nil {
		panic(err)
	}
	if err := connection.Provide(container); err != nil {
		panic(err)
	}
	if err := repository.Provide(container); err != nil {
		panic(err)
	}
	if err := service.Provide(container); err != nil {
		panic(err)
	}
	if err := controller.Provide(container); err != nil {
		panic(err)
	}
	e := echo.New()
	e.Use(custom_middleware.ZeroLog(log.Logger))
	e.Use(middleware.RequestID())
	e.Use(middleware.Recover())
	g := e.Group("/api")
	if err := controller.InitRouter(container, g); err != nil {
		panic(err)
	}
	e.Logger.Fatal(e.Start(":8080"))

}
