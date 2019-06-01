package main

import (
  "github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gin-sample/driver"
  "gin-sample/router"
  "strconv"
)

func main() {
  port := 80

  // Load ENV
	if err := godotenv.Load(); err != nil {
    panic(".env file cannot be load.")
	}

  // DB connection
  driver.Init()
  defer driver.DB.Close()

  // Routing
  r := gin.Default()
  router.App(r)
  r.Run(":" + strconv.Itoa(port))
}
