package router

import (
	. "gin-sample/http/controller"
	"github.com/gin-gonic/gin"
)

func App(router *gin.Engine) {
	version := "/v1"
	router.GET(version+"/designs", DesignIndex)
}
