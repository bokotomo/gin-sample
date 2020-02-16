package router

import (
	"gin-sample/http/controller"
	"gin-sample/http/middleware"

	"github.com/gin-gonic/gin"
)

// App is
func App(router *gin.Engine) {
	version := "/v1"
	v1 := router.Group(version)

	// None Authentication
	{
		// Auth
		v1.POST("/login", controller.LoginIndex)

		// User
		v1.POST("/user", controller.UserStore)
	}

	// Authentication
	auth := v1.Group("", middleware.AuthMiddleware())
	{
		// Auth
		auth.GET("/authorize", controller.AuthorizeIndex)

		// Designs
		auth.GET("/designs", controller.DesignIndex)
		auth.GET("/design/:designId", controller.DesignShow)
	}
}
