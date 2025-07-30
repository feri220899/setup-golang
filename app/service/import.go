package service

import (
	"encoding/base64"
	"encoding/csv"
	"fmt"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
)

func CekTemplate(filepath string, modul string, request *gin.Context) int {
	path_decode, _ := base64.StdEncoding.DecodeString(filepath)
	switch modul {
	case "test":
		return TestImport(string(path_decode), request)
	case "modul2":
		return modul2Import(string(path_decode), request)
	case "asersi_cvc":
		return asersiCvcImport(string(path_decode), request)
	default:
		request.AbortWithStatusJSON(400, gin.H{
			"error": fmt.Sprintf("Modul %s tidak ditemukan", modul),
		})
		return 0
	}
}

func TestImport(filepath string, request *gin.Context) int {
	file_read, _ := excelize.OpenFile(filepath)
	rows, _ := file_read.GetRows("Sheet1")
	if len(rows) < 2 {
		request.AbortWithStatusJSON(400, gin.H{
			"error": "File tidak berisi data",
		})
		return 0
	}
	if len(rows[0]) < 9 {
		request.AbortWithStatusJSON(400, gin.H{
			"error": "File tidak sesuai template",
		})
		return 0
	}
	expectedHeader := map[int]string{
		0: "Nama_Kolom1",
		1: "Nama_Kolom2",
		2: "Nama_Kolom3",
		3: "Nama_Kolom4",
		4: "Nama_Kolom5",
		5: "Nama_Kolom6",
		6: "Nama_Kolom7",
		7: "Nama_Kolom8",
		8: "Nama_Kolom9",
	}
	for i, expected := range expectedHeader {
		if rows[0][i] != expected {
			request.AbortWithStatusJSON(400, gin.H{
				"error": fmt.Sprintf("Template Tidak Sesuai Kolom ke-%d harus '%s'", i+1, expected),
			})
			return 0
		}
	}
	totalRows := len(rows) - 1
	return totalRows
}

func modul2Import(filepath string, request *gin.Context) int {
	file_read, _ := excelize.OpenFile(filepath)
	rows, _ := file_read.GetRows("Sheet1")
	if len(rows) < 2 {
		request.AbortWithStatusJSON(400, gin.H{"error": "File is empty or does not contain data"})
		return 0
	}
	if len(rows[0]) < 9 {
		request.AbortWithStatusJSON(400, gin.H{"error": "File does not match the template"})
		return 0
	}
	expectedHeader := map[int]string{
		0: "Nama_Kolom1",
		1: "Nama_Kolom2",
		2: "Nama_Kolom3",
		3: "Nama_Kolom4",
		4: "Nama_Kolom5",
		5: "Nama_Kolom6",
		6: "Nama_Kolom7",
		7: "Nama_Kolom8",
		8: "Nama_Kolom9",
	}
	if len(rows) == 0 || len(rows[0]) < len(expectedHeader) {
		request.AbortWithStatusJSON(400, gin.H{"error": "Header tidak lengkap"})
		return 0
	}
	for i, expected := range expectedHeader {
		if rows[0][i] != expected {
			request.AbortWithStatusJSON(400, gin.H{
				"error": fmt.Sprintf("Kolom ke-%d harus '%s'", i+1, expected),
			})
			return 0
		}
	}
	totalRows := len(rows) - 1
	return totalRows
}

func asersiCvcImport(filepath string, request *gin.Context) int {
	file, _ := os.Open(filepath)
	defer file.Close()
	reader := csv.NewReader(file)
	reader.Comma = ';'
	header, _ := reader.Read()
	rows, _ := reader.ReadAll()
	for i := range header {
		header[i] = strings.TrimSpace(strings.ToLower(header[i]))
	}
	if len(header) < 21 {
		request.AbortWithStatusJSON(400, gin.H{
			"error": "File tidak sesuai template",
		})
		return 0
	}
	expectedHeader := map[int]string{
		0:  "hose_id",
		1:  "hose_number",
		2:  "pump_id",
		3:  "pump_name",
		4:  "attendant_name",
		5:  "delivery_id",
		6:  "site_id",
		7:  "city",
		8:  "address",
		9:  "product",
		10: "completed_date",
		11: "jam",
		12: "delivery_type",
		13: "del_sell_price",
		14: "delivery_volume",
		15: "delivery_value",
		16: "vehicle_number",
		17: "keterangan",
		18: "batch",
		19: "sector",
		20: "nik",
	}
	for i, expected := range expectedHeader {
		if header[i] != expected {
			request.AbortWithStatusJSON(400, gin.H{
				"error": fmt.Sprintf("Template Tidak Sesuai Kolom ke-%d harus '%s'", i+1, expected),
			})
			return 0
		}
	}
	totalRows := len(rows) - 1
	return totalRows
}
