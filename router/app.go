package router

import (
	"github.com/gin-gonic/gin"
	"softwareCup/service"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", service.GetMessage)
	return r
}
