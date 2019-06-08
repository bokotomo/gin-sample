package model

import (
	"github.com/jinzhu/gorm"
)

type Design struct {
	gorm.Model
	Title     string
	Text      string
	Good      int
	Comments  int
	Thumbnail string
}
