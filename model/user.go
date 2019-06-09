package model

type User struct {
	ID       uint   `gorm:"primary_key"`
	Password string `gorm:"not null"`
	Email    string `gorm:"not null"`
}
