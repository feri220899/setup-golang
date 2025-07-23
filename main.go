package main

import (
	"fmt"
	"golang-restfull-api/app/helper"
	config "golang-restfull-api/config"
	routes "golang-restfull-api/router"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println(helper.HashPassword("password2"))
	db := config.ConnectDB()
	router := gin.Default()
	routes.UserRoutes(router, db)
	config.StartServer(router)
}
