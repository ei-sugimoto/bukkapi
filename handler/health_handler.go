package handler

import (
	"net/http"

	"github.com/ei-sugimoto/bukkapi/handler/views"
	"github.com/labstack/echo/v4"
)

func HeathHandler(c echo.Context) error {
	res := &views.HealthView{
		Status: "healthy",
	}

	return c.JSON(http.StatusOK, res)
}
