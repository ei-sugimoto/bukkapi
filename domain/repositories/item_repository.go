package repositories

import "github.com/ei-sugimoto/bukkapi/domain/entities"

type ItemRepository interface {
	List() (*entities.Item, error)
	Create(name string) (*entities.Item, error)
}
