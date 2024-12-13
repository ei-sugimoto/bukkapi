package usecases

import (
	"context"

	"github.com/ei-sugimoto/bukkapi/domain/entities"
	"github.com/ei-sugimoto/bukkapi/domain/repositories"
)

type CreateItemInput struct {
	Name string
}

type ItemUsecase struct {
	ItemRepo repositories.ItemRepository
}

func NewItemUsecase(repo repositories.ItemRepository) *ItemUsecase {
	return &ItemUsecase{
		ItemRepo: repo,
	}
}

func (i *ItemUsecase) Create(c context.Context, input CreateItemInput) error {
	newItem := &entities.Item{
		Name: input.Name,
	}

	if err := i.ItemRepo.Create(c, newItem); err != nil {
		return err
	}

	return nil
}
