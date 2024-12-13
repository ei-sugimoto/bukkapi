package models

import "gorm.io/gorm"

type Item struct {
	gorm.Model
	Name string
}

// TableName sets the insert table name for this struct type
func (Item) TableName() string {
	return "items"
}
