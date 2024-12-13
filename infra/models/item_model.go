package models

type ItemModel struct {
	ID   int    `gorm:"primaryKey;autoIncrement"`
	Name string `gorm:"size:100;not null"`
}
