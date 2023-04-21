package service

import (
	"github.com/gin-gonic/gin"
	"softwareCup/utils"
)

func GetMessage(c *gin.Context) {
	sql := "insert into file values (?,?)"
	utils.Db.Exec(sql, 1, "path")
	c.JSON(200, gin.H{
		"message": "succeed",
	})
}
