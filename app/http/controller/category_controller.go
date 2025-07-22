package controller

import (
	"encoding/base64"
	"fmt"
	"golang-restfull-api/app/helper"
	categorymodel "golang-restfull-api/app/model/category"
	"golang-restfull-api/app/model/import/testmodule"
	"golang-restfull-api/app/service"
	http "net/http"
	"os"
	"path/filepath"
	"time"

	gin "github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
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
	file_name := filepath.Ext(file.Filename)
	path := helper.PublicPath("uploads", fmt.Sprintf("data_%d%s", time.Now().Unix(), file_name))
	request.SaveUploadedFile(file, path)
	path_encode := base64.StdEncoding.EncodeToString([]byte(path))

	// CEK TEMPLATE
	totalRows := service.CekTemplate(path_encode, "test", request)
	if totalRows == 0 {
		path_decode, _ := base64.StdEncoding.DecodeString(path_encode)
		os.Remove(string(path_decode))
		return
	}

	// Simpan ke database
	import_status := testmodule.TestModule{
		Import_file_path: path_encode,
		Import_status:    "processing",
		Import_start:     1,
		Import_batch:     1000,
		Import_total_row: totalRows,
	}
	if create_import_status := db.Table("import_status").Create(&import_status).Error; create_import_status != nil {
		request.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to save import status",
		})
		os.Remove(path)
		return
	}

	request.JSON(http.StatusOK, gin.H{
		"message":          "File uploaded successfully",
		"import_status_id": import_status.Id,
	})
}

func ImportProgres(request *gin.Context, db *gorm.DB) {
	import_status_id := request.PostForm("import_status_id")
	var import_status testmodule.TestModule
	if err := db.Table("import_status").Where("id", import_status_id).First(&import_status).Error; err != nil {
		request.JSON(http.StatusNotFound, gin.H{
			"error": "Import status not found",
		})
		return
	}
	if import_status.Import_status == "processing" {
		if import_status.Import_start >= import_status.Import_total_row {
			db.Table("import_status").Where("id", import_status_id).Updates(&testmodule.TestModule{
				Import_status: "completed",
			})
			request.JSON(http.StatusOK, gin.H{
				"message":      "Import process completed",
				"Import_start": import_status.Import_start,
			})
			return
		} else {
			path_decode, _ := base64.StdEncoding.DecodeString(import_status.Import_file_path)
			start := import_status.Import_start
			batch := import_status.Import_batch
			file_read, _ := excelize.OpenFile(string(path_decode))

			var data []testmodule.DataExcel
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
				data = append(data, testmodule.DataExcel{
					Nama_Kolom1: row[0],
					Nama_Kolom2: row[1],
					Nama_Kolom3: row[2],
					Nama_Kolom4: row[3],
					Nama_Kolom5: row[4],
					Nama_Kolom6: row[5],
					Nama_Kolom7: row[6],
					Nama_Kolom8: row[7],
					Nama_Kolom9: row[8],
				})
				numrow++
			}
			new_start := min(start+batch, import_status.Import_total_row)
			db.Table("import_status").Where("id", import_status_id).Updates(&testmodule.TestModule{
				Import_start: new_start,
			})

			db.Table("data_excel").CreateInBatches(data, 500)

			request.JSON(http.StatusOK, gin.H{
				"message":      "Import process in progress",
				"data":         data,
				"Import_start": new_start,
			})
		}
	} else if import_status.Import_status == "paused" {
		request.JSON(http.StatusOK, gin.H{
			"message":      "Import process is paused",
			"data":         import_status,
			"Import_start": import_status.Import_start,
		})
	}
	if import_status.Import_status == "completed" {
		request.JSON(http.StatusOK, gin.H{
			"message":      "Import process completed",
			"data":         import_status,
			"Import_start": import_status.Import_start,
		})
	}
}
