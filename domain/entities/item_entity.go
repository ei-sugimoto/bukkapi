package entities

import "errors"

type Item struct {
	ID   int
	Name string
}

var ErrNameEmpty = errors.New("name empty")

func NewItem(id int, name string) (*Item, error) {
	if name == "" {
		return nil, ErrNameEmpty
	}

	return &Item{
		ID:   id,
		Name: name,
	}, nil
}
