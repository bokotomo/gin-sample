package main

import (
	. "gin-sample/driver"
	. "gin-sample/model"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		panic(".env file cannot be load.")
	}
	Init()
	DB.AutoMigrate(&Design{}, &Token{}, &User{})
	defer DB.Close()
}
