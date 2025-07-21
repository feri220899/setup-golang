package controller

import (
	"encoding/base64"
	"fmt"
	"golang-restfull-api/app/helper"
	categorymodel "golang-restfull-api/app/model/category"
	http "net/http"
	"path/filepath"
	"time"

	gin "github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
	gorm "gorm.io/gorm"
)

type DataExcel struct {
	Row1 string
	Row2 string
	Row3 string
	Row4 string
	Row5 string
	Row6 string
	Row7 string
	Row8 string
	Row9 string
}

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
	file_name := filepath.Ext(file.Filename)
	path := helper.PublicPath("uploads", fmt.Sprintf("data_%d%s", time.Now().Unix(), file_name))
	request.SaveUploadedFile(file, path)
	// hitung total row di mulai dari baris ke-2
	file_read, _ := excelize.OpenFile(path)
	rows, _ := file_read.GetRows("Sheet1")
	totalRows := len(rows) - 1

	path_encode := base64.StdEncoding.EncodeToString([]byte(path))
	request.JSON(http.StatusOK, gin.H{
		"message": "File uploaded successfully",
		"path":    path_encode,
		"start":   1,
		"batch":   5,
		"total":   totalRows,
	})
}

func ImportProgres(request *gin.Context, db *gorm.DB) {
	path := request.PostForm("path")
	path_decode, _ := base64.StdEncoding.DecodeString(path)
	start := 1 // baris ke-2 ini bisa diambil dari parameter request
	batch := 5 // jumlah data yang diambil per batch ini bisa diambil dari parameter request
	file_read, _ := excelize.OpenFile(string(path_decode))

	var data []DataExcel
	rows, _ := file_read.GetRows("Sheet1")
	numrow := 0
	for _, row := range rows {
		if numrow < start {
			numrow++
			continue
		}
		if numrow >= start+batch {
			break
		}
		data = append(data, DataExcel{
			Row1: row[0],
			Row2: row[1],
			Row3: row[2],
			Row4: row[3],
			Row5: row[4],
			Row6: row[5],
			Row7: row[6],
			Row8: row[7],
			Row9: row[8],
		})
		numrow++
	}
	new_start := start + batch
	helper.DumpToBrowser(request,
		"data", data,
		"new_start", new_start,
	)
}
