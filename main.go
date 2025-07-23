package main

import (
	config "golang-restfull-api/config"
	routes "golang-restfull-api/router"

	"github.com/gin-gonic/gin"
)

func main() {
	db := config.ConnectDB()
	router := gin.Default()
	routes.UserRoutes(router, db)
	config.StartServer(router)
}
