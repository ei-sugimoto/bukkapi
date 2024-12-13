package repositories

import (
	"context"

	"github.com/ei-sugimoto/bukkapi/domain/entities"
)

type ItemRepository interface {
	List(c context.Context) (*[]entities.Item, error)
	Create(c context.Context, item *entities.Item) error
}
