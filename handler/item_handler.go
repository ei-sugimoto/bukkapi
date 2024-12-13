package handler

import (
	"net/http"

	"github.com/ei-sugimoto/bukkapi/app/usecases"
	"github.com/labstack/echo/v4"
)

type ItemHandler struct {
	ItemUsecase *usecases.ItemUsecase
}

func NewItemHandler(uc *usecases.ItemUsecase) *ItemHandler {
	return &ItemHandler{
		ItemUsecase: uc,
	}
}

type createReq struct {
	Name string `json:"name"`
}

func (ih *ItemHandler) Create(c echo.Context) error {
	req := new(createReq)
	if err := c.Bind(req); err != nil {
		return err
	}

	if req.Name == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Name is required"})
	}
	input := usecases.CreateItemInput{
		Name: req.Name,
	}

	err := ih.ItemUsecase.Create(c.Request().Context(), input)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, "")
}
