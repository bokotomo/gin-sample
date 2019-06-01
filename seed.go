package main

import (
  "github.com/jinzhu/gorm"
	. "gin-sample/driver"
	. "gin-sample/model"
)

func main() {
  Init()
  DB.Create(&Design{Model: gorm.Model{ID: 1}, Title: "design1"})
  DB.Create(&Design{Model: gorm.Model{ID: 2}, Title: "design2"})
  DB.Create(&Design{Model: gorm.Model{ID: 3}, Title: "design3"})
  defer DB.Close()
}