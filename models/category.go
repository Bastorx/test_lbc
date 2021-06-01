package models

import (
	"gorm.io/gorm"
	"time"
)

type Category struct {
	ID        uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Name  string    `gorm:"size:100;not null;" json:"name"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (c *Category) Create(db *gorm.DB) (*Category, error) {
	result := db.Create(c)
}