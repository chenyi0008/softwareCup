package router

import (
	"github.com/gin-gonic/gin"
	"softwareCup/service"
	"softwareCup/utils"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.Use(utils.Cors())
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
		files.GET("", service.GetList)
	}

	return r
}
