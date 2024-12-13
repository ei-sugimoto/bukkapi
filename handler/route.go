package handler

import (
	"github.com/ei-sugimoto/bukkapi/app/usecases"
	"github.com/ei-sugimoto/bukkapi/infra/database"
	"github.com/ei-sugimoto/bukkapi/infra/repositories"
	"github.com/labstack/echo/v4"
)

func SetRoute(e *echo.Echo) {
	e.GET("/health", HeathHandler)

	ItemHandler, err := CreateItemHandler()
	if err != nil {
		e.Logger.Fatal(err)
	}

	e.POST("/items", ItemHandler.Create)

}

func CreateItemHandler() (*ItemHandler, error) {
	db, err := database.Connect()
	if err != nil {
		return nil, err
	}

	repo := repositories.NewItemRepository(db)
	uc := usecases.NewItemUsecase(repo)
	ha := NewItemHandler(uc)

	return ha, nil
}
