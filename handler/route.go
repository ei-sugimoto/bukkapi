package handler

import (
	"github.com/labstack/echo/v4"
)

func SetRoute(e *echo.Echo) {
	e.GET("/health", HeathHandler)

}
