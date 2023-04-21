package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"softwareCup/models"
	"softwareCup/utils"
	"strconv"
)

func GetMessage(c *gin.Context) {
	sql := "insert into file values (?,?)"
	utils.Db.Exec(sql, 1, "path")

	c.JSON(200, gin.H{
		"message": "succeed",
	})
}

func UploadOne(c *gin.Context) {

	id, _ := utils.GetID()
	fmt.Println("id:")
	fmt.Println(id)
	file, _ := c.FormFile("file")

	model := models.FileModel{Path: strconv.Itoa(int(id)) + "_" + file.Filename}
	c.SaveUploadedFile(file, viper.GetString("file.path")+strconv.Itoa(int(id))+"_"+file.Filename)
	_, err := utils.Db.Insert(model)
	if err != nil {
		fmt.Println(err)
	}
}

func GetOne(c *gin.Context) {
	id := c.Query("fileName")
	fmt.Println(id)
	c.File(viper.GetString("file.path") + id)
}

func UploadSome(c *gin.Context) {
	form, _ := c.MultipartForm()
	fileList := form.File["files"]

	for _, file := range fileList {
		id, _ := utils.GetID()
		model := models.FileModel{Path: strconv.Itoa(int(id)) + "_" + file.Filename}
		c.SaveUploadedFile(file, viper.GetString("file.path")+strconv.Itoa(int(id))+"_"+file.Filename)
		_, err := utils.Db.Insert(model)
		if err != nil {
			fmt.Println(err)
		}
	}
}
