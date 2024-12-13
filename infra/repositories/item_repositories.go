package repositories

import (
	"context"

	"github.com/ei-sugimoto/bukkapi/domain/entities"
	"github.com/ei-sugimoto/bukkapi/domain/repositories"
	"github.com/ei-sugimoto/bukkapi/infra/models"
	"gorm.io/gorm"
)

type ItemRepositoryImp struct {
	DB *gorm.DB
}

func NewItemRepository(db *gorm.DB) repositories.ItemRepository {
	return &ItemRepositoryImp{DB: db}
}

func (i *ItemRepositoryImp) List(c context.Context) (*[]entities.Item, error) {
	var items []models.Item

	if err := i.DB.Find(&items); err != nil {
		return nil, err.Error
	}
	return nil, nil
}

func (i *ItemRepositoryImp) Create(c context.Context, item *entities.Item) error {
	newItem := i.formEntity(item)
	err := i.DB.Save(newItem)
	if err != nil {
		return err.Error
	}

	return nil
}

func (i *ItemRepositoryImp) formEntity(e *entities.Item) *models.Item {
	return &models.Item{
		Name: e.Name,
	}
}
