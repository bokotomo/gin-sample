package router

import (
	command "gin-sample/command/http/controller"
	query "gin-sample/query/http/controller"
	"gin-sample/query/http/middleware"

	"github.com/gin-gonic/gin"
)

// App is
func App(router *gin.Engine) {
	version := "/v1"
	v1 := router.Group(version)

	// None Authentication
	{
		// Auth
		v1.POST("/login", command.LoginIndex)

		// User
		v1.POST("/user", command.UserStore)
	}

	// Authentication
	auth := v1.Group("", middleware.AuthMiddleware())
	{
		// Auth
		auth.GET("/authorize", command.AuthorizeIndex)

		// Designs
		auth.GET("/designs", query.DesignIndex)
		auth.GET("/design/:designId", query.DesignShow)
	}
}
