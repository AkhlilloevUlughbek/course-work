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

var db *gorm.DB

func main() {
	var err error
	db, err = gorm.Open(postgres.Open("host=localhost port=5432 user=scientist password=scientist database=researches sslmode=disable"), &gorm.Config{
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
	router.PUT("/users/:otp", ConfirmOTP)
	router.GET("/users/:email", ResetPassword)
	router.PUT("/users/reset-password/:email", FinishReset)
	router.POST("/login", Login)

	err = http.ListenAndServe("localhost:8080", router)
	if err != nil {
		log.Fatalf("error: %v", err.Error())
	}
}
