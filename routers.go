package main

import (
	"acussm/demo/controller"
	"acussm/demo/middlleware"
	"github.com/gin-gonic/gin"
)

func collectRoutE(r *gin.Engine) *gin.Engine {
	r.POST("/api/auth/register", controller.Register)
	r.POST("/api/auth/login",controller.Login)
	r.GET("/api/auth/info",middlleware.Authmiddleware(),controller.Info)

	return r
}