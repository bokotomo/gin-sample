package main

import (
	. "gin-sample/driver"
	. "gin-sample/model"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(".env file cannot be load.")
	}
	TestInit()
	TestDB.AutoMigrate(&Design{})
	defer TestDB.Close()
}
