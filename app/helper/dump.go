package helper

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func DumpContextJSON(c *gin.Context) {
	var bodyData interface{}
	contentType := c.GetHeader("Content-Type")

	// Backup dan baca raw body (hanya untuk keperluan reset)
	if c.Request.Body != nil {
		bodyBytes, err := io.ReadAll(c.Request.Body)
		if err == nil {
			// Reset supaya Gin bisa membaca lagi
			c.Request.Body = io.NopCloser(bytes.NewReader(bodyBytes))
		}
	}

	// Handle sesuai Content-Type
	switch {
	case strings.Contains(contentType, "application/json"):
		var jsonData map[string]interface{}
		if err := c.ShouldBindJSON(&jsonData); err == nil {
			bodyData = jsonData
		}
	case strings.Contains(contentType, "application/x-www-form-urlencoded"):
		_ = c.Request.ParseForm()
		formData := map[string][]string(c.Request.PostForm)
		bodyData = formData
	case strings.Contains(contentType, "multipart/form-data"):
		if err := c.Request.ParseMultipartForm(32 << 20); err == nil {
			form := make(map[string]interface{})
			for key, values := range c.Request.MultipartForm.Value {
				if len(values) == 1 {
					form[key] = values[0]
				} else {
					form[key] = values
				}
			}
			bodyData = form
		}
	default:
		bodyData = "Unsupported Content-Type: " + contentType
	}

	data := map[string]interface{}{
		"method":  c.Request.Method,
		"path":    c.Request.URL.Path,
		"headers": filterCustomHeaders(c.Request.Header),
		"query":   c.Request.URL.Query(),
		"params":  c.Params,
		"body":    bodyData,
	}

	c.JSON(http.StatusOK, data)
	c.Abort()
}

func filterCustomHeaders(allHeaders http.Header) map[string][]string {
	ignored := map[string]bool{
		"Accept":          true,
		"Accept-Encoding": true,
		"Connection":      true,
		"Content-Length":  true,
		"Content-Type":    true,
		"User-Agent":      true,
		"Postman-Token":   true,
		"Cache-Control":   true,
		"Host":            true,
	}

	filtered := make(map[string][]string)
	for key, value := range allHeaders {
		if !ignored[key] && !strings.HasPrefix(strings.ToLower(key), "accept") {
			filtered[key] = value
		}
	}
	return filtered
}

func DumpToBrowser(c *gin.Context, data ...interface{}) {
	typeDump := []map[string]interface{}{}

	for _, d := range data {
		typeDump = append(typeDump, map[string]interface{}{
			"type":  fmt.Sprintf("%T", d),
			"value": d,
		})
	}

	c.Header("Content-Type", "application/json")
	c.Status(http.StatusOK)

	jsonData, err := json.MarshalIndent(typeDump, "", "  ")
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("Dump error: %v", err))
		return
	}
	c.Writer.Write(jsonData)
	c.Abort()
}

func AbortWithJSON(c *gin.Context, statusCode int, status string, message string, data interface{}) {
	response := gin.H{
		"status":  status,
		"message": message,
	}
	if data != nil {
		response["data"] = data
	}

	c.AbortWithStatusJSON(statusCode, response)
}
