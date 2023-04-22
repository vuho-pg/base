package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/vuho-pg/base/pkg/util/api"
	"go.uber.org/dig"
)

func InitRouter(c *dig.Container, e *echo.Group) error {
	if err := c.Invoke(func(
		demoController DemoController,
	) {
		cons := []api.EchoController{
			demoController,
		}
		for _, con := range cons {
			con.Register(e)
		}
	}); err != nil {
		return err
	}
	return nil
}
