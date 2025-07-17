package main

import (
	config "golang-restfull-api/config"
	routes "golang-restfull-api/router"

	"github.com/gin-gonic/gin"
)

func main() {
	db := config.ConnectDB()
	r := gin.Default()
	routes.UserRoutes(r, db)
	r.Run(":8080")
}
