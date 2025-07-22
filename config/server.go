package config

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func StartServer(router *gin.Engine) {
	app_url := viper.GetString("APP_URL")
	app_port := viper.GetString("APP_PORT")
	time_out := viper.GetInt("APP_REQUEST_TIMEOUT")

	s := &http.Server{
		Addr:           ":" + app_port,
		Handler:        router,
		ReadTimeout:    time.Duration(time_out) * time.Second,
		WriteTimeout:   time.Duration(time_out) * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	fmt.Printf("Server running at %s:%s", app_url, app_port)
	fmt.Println(" with a timeout of", time_out, "seconds")
	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("ListenAndServe error: %v", err)
	}
	s.ListenAndServe()
}
