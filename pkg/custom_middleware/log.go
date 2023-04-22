package custom_middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
)

func ZeroLog(l zerolog.Logger) echo.MiddlewareFunc {
	return middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		Skipper: nil,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			l.Info().
				Str("method", v.Method).
				Str("request_id", v.RequestID).
				Str("host", v.Host).
				Str("uri", v.URI).
				Int("status", v.Status).
				Err(v.Error).
				Str("user_agent", v.UserAgent).
				Msg("request")
			return nil
		},
		LogHost:      true,
		LogMethod:    true,
		LogURI:       true,
		LogURIPath:   true,
		LogRoutePath: true,
		LogRequestID: true,
		LogUserAgent: true,
		LogStatus:    true,
		LogError:     true,
		HandleError:  true,
	})
}
