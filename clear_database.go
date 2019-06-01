package main

import (
	. "gin-sample/driver"
	. "gin-sample/model"
)

func main() {
  Init()
  DB.DropTable(&Design{})
  defer DB.Close()
}
