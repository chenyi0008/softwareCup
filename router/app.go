package router

import (
	"github.com/gin-gonic/gin"
	"softwareCup/service"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", service.GetMessage)

	users := r.Group("/users")
	{
		users.POST("/register", service.Register)
		users.POST("/login", service.Login)
	}

	files := r.Group("/files")
	{
		files.POST("/uploadOne", service.UploadOne)
		files.GET("/getOne", service.GetOne)
		files.POST("/uploadSome", service.UploadSome)
	}

	return r
}
