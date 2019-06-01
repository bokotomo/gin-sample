package main

import (
	. "gin-sample/driver"
	. "gin-sample/model"
)

func main() {
  Init()
  DB.AutoMigrate(&Design{})
  defer DB.Close()
}
