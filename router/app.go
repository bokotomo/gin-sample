package router

import (
  "github.com/gin-gonic/gin"
  "gin-sample/http/controller"
)

func App(router *gin.Engine) {
  router.GET("/designs", controller.DesignIndex)
}
