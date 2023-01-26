package main

import (
	"gin-sample/driver"
	"gin-sample/model"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		panic(".env file cannot be load.")
	}
	driver.Init()
	driver.DB.DropTable(&model.Design{}, &model.Token{}, &model.User{})
	defer driver.DB.Close()
}
