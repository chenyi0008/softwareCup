package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"softwareCup/models"
	"softwareCup/utils"
)

func Register(c *gin.Context) {

	username := c.PostForm("username")
	password := c.PostForm("password")

	var tmp models.User
	utils.Db.Where("username = ?", username).Get(&tmp)
	fmt.Println(tmp.Id)
	if tmp.Id != 0 {
		c.JSON(200, gin.H{
			"code": 0,
			"msg":  "此账号已注册",
			"data": nil,
		})
		return
	}

	user := models.User{
		Username: username,
		Password: password,
	}

	utils.Db.Insert(user)

	c.JSON(200, gin.H{
		"code": 1,
		"msg":  "注册成功",
		"data": nil,
	})
}

func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	var tmp models.User
	utils.Db.Where("username = ? and password = ?", username, password).Get(&tmp)
	if tmp.Id == 0 {
		c.JSON(200, gin.H{
			"code": 0,
			"msg":  "账号或密码输入有误",
			"data": nil,
		})
		return
	}

	c.JSON(200, gin.H{
		"code": 1,
		"msg":  "登录成功",
		"data": nil,
	})
}
