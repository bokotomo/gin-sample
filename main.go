package main

import (
  "github.com/gin-gonic/gin"
  "strconv"
  "gin-sample/router"
)

func main() {
  port := 80

  r := gin.Default()
  router.App(r)
  r.Run(":" + strconv.Itoa(port))
}
