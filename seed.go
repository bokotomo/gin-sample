package main

import (
	. "gin-sample/driver"
	. "gin-sample/model"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(".env file cannot be load.")
	}
	Init()
	DB.Unscoped().Delete(&Design{})
	DB.Create(&Design{Model: gorm.Model{ID: 1}, Title: "design1"})
	DB.Create(&Design{Model: gorm.Model{ID: 2}, Title: "design2"})
	DB.Create(&Design{Model: gorm.Model{ID: 3}, Title: "design3"})
	DB.Create(&Design{Model: gorm.Model{ID: 4}, Title: "design4"})
	DB.Create(&Design{Model: gorm.Model{ID: 5}, Title: "design5"})
	DB.Create(&Design{Model: gorm.Model{ID: 6}, Title: "design6"})
	DB.Create(&Design{Model: gorm.Model{ID: 7}, Title: "design7"})
	DB.Create(&Design{Model: gorm.Model{ID: 8}, Title: "design8"})
	defer DB.Close()
}
