package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"net/http"
	"os"
	"time"
)

var DB *gorm.DB

func main() {
	DB, err := gorm.Open(postgres.Open("host=%s port=%d user=%s password=%s database=%s sslmode=%s"), &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				SlowThreshold:             time.Second,
				LogLevel:                  logger.Info,
				IgnoreRecordNotFoundError: true,
				Colorful:                  false,
			},
		),
	})

	if err != nil {
		log.Fatal("failed to init DB", err)
	}
	router := gin.New()
	router.Use(gin.Recovery())
	router.POST("/users", CreateUser)

	err = http.ListenAndServe("localhost:8080", router)
	if err != nil {
		log.Fatalf("error: %v", err.Error())
	}
}
