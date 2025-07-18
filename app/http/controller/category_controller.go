package controller

import (
	categorymodel "golang-restfull-api/app/model/category"
	http "net/http"

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
