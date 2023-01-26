package main

import (
	"gin-sample/driver"
	"gin-sample/router"
	"gin-sample/util"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// run is
func run(port int) {
	// Load ENV
	if err := godotenv.Load(); err != nil {
		panic(".env file cannot be load.")
	}

	// DB connection
	driver.Init()
	defer driver.DB.Close()

	// Middleware init
	util.FileInit()

	// Routing
	r := gin.Default()
	router.App(r)
	r.Run(":" + strconv.Itoa(port))
}

func main() {
	port := 80
	run(port)
}
