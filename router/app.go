package router

import (
	. "gin-sample/http/controller"
	"github.com/gin-gonic/gin"
	// jwt "github.com/dgrijalva/jwt-go"
	// "github.com/dgrijalva/jwt-go/request"
)

// var secretKey = "75c92a074c341e9964329c0550c2673730ed8479c885c43122c90a2843177d5ef21cb50cfadcccb20aeb730487c11e09ee4dbbb02387242ef264e74cbee97213"

func App(router *gin.Engine) {
	version := "/v1"
	router.GET(version+"/designs", DesignIndex)
	router.GET(version+"/design/:designId", DesignShow)
}
