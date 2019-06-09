package model

import (
	// "github.com/jinzhu/gorm"
	"time"
)

type Token struct {
	ID     uint      `gorm:"primary_key"`
	UserID uint      `gorm:"not null"`
	Token  string    `gorm:"not null;size:440"`
	Exp    time.Time `gorm:"not null;type:datetime"`
}
