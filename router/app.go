package router

import (
  "github.com/gin-gonic/gin"
  . "gin-sample/http/controller"
)

func App(router *gin.Engine) {
  version := "/v1"
  router.GET(version+"/designs", DesignIndex)
}
