package controller

import (
	"fmt"
	"golang-restfull-api/app/helper"
	categorymodel "golang-restfull-api/app/model/category"
	http "net/http"
	"path/filepath"
	"time"

	gin "github.com/gin-gonic/gin"
	gorm "gorm.io/gorm"
)

func GetUsers(request *gin.Context, db *gorm.DB) {
	var results []categorymodel.Category
	db.Table("category").Select("id, name, phone_number").Find(&results)
	request.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    results,
	})
}

func ImportData(request *gin.Context, db *gorm.DB) {
	file, _ := request.FormFile("file")
	fmt.Println("File received:", file.Filename)
	file_name := filepath.Ext(file.Filename)
	path := helper.PublicPath("uploads", fmt.Sprintf("data_%d%s", time.Now().Unix(), file_name))
	request.SaveUploadedFile(file, path)
}
