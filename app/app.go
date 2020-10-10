package app

import "github.com/labstack/echo/v4"

func Router() func(e *echo.Echo) {
	return func(e *echo.Echo) {
		e.GET("/ping", func(context echo.Context) error {
			return context.JSON(200, "pong")
		})
	}
}
