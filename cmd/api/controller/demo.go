package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/vuho-pg/base/pkg/service"
	"github.com/vuho-pg/base/pkg/util/api"
)

type DemoController interface {
	api.EchoController
}

type demoController struct {
	demoService service.DemoService
}

func NewDemoController(demoService service.DemoService) DemoController {
	return &demoController{
		demoService: demoService,
	}
}

func (d *demoController) Register(g *echo.Group) {
	g = g.Group("/demo")
	g.POST("", d.CreateDemo)
}

func (d *demoController) CreateDemo(c echo.Context) error {
	return api.ServeEcho(c, d.demoService.Create)
}
