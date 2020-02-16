package model

import (
	"github.com/jinzhu/gorm"
)

// Design is
type Design struct {
	gorm.Model
	Title     string `gorm:"not null"`
	Text      string `gorm:"not null"`
	Good      int    `gorm:"not null"`
	Comments  int    `gorm:"not null"`
	Thumbnail string `gorm:"not null"`
}
