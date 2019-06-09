package router

import (
	. "gin-sample/http/controller"
	. "gin-sample/http/middleware"
	"github.com/gin-gonic/gin"
)

func App(router *gin.Engine) {
	version := "/v1"
	v1 := router.Group(version)

	// None Authentication
	{
		// Auth
		v1.POST("/login", LoginIndex)

		// User
		v1.POST("/user", UserStore)
	}

	// Authentication
	auth := v1.Group("", AuthMiddleware())
	{
		// Auth
		auth.GET("/authorize", AuthorizeIndex)

		// Designs
		auth.GET("/designs", DesignIndex)
		auth.GET("/design/:designId", DesignShow)
	}
}
