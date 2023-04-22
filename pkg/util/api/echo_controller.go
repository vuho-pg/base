package api

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/vuho-pg/base/pkg/util/errors"
)

var globalValidator = validator.New()

type EchoController interface {
	Register(e *echo.Group)
}

type echoController struct {
}

func (echoController) ServeError(e echo.Context, err error) error {
	resp := FromError(err)
	return e.JSON(resp.GetCode(), resp)
}

func (echoController) ServeResponse(e echo.Context, resp Response) error {
	return e.JSON(resp.GetCode(), resp)
}

func ServeEcho[T any](e echo.Context, fn func(ctx context.Context, data T) (Response, error)) error {
	var c echoController
	var data T
	if err := e.Bind(&data); err != nil {
		return c.ServeError(e, err)
	}
	if err := globalValidator.Struct(data); err != nil {
		return c.ServeError(e, errors.Wrap(err, errors.ErrBadRequest))
	}
	resp, err := fn(e.Request().Context(), data)
	if err != nil {
		return c.ServeError(e, err)
	}
	return c.ServeResponse(e, resp)
}
