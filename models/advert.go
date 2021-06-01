package models

import (
	"time"
	"gorm.io/datatypes"
)

type Advert struct {
	ID        uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Title  string    `gorm:"size:100;not null;" json:"name"`
	Content  string    `gorm:"size:100;not null;" json:"name"`
	CategoryID uint32
	Category Category `gorm:"foreignKey:CategoryID"`
	Payload datatypes.JSON
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}