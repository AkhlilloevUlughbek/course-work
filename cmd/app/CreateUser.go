package main

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"math/rand"
	"time"
)

func CreateUser(c *gin.Context) {
	var newUser User
	err := c.ShouldBindJSON(&newUser)
	if err != nil {
		c.JSONP(400, errors.New(fmt.Sprintf("err: %v", err.Error())))
		return
	}

	if !userValidation(newUser) {
		c.JSONP(400, errors.New("you forgot to sent some relevant field(s). please, check everything and try again"))
	}

	otp := generateSixDigitCode()

	err = saveCodeToDB(newUser.Email, otp)
	if err != nil {
		log.Printf("failed to save user: %v", err)
		c.JSONP(500, errors.New(fmt.Sprintf("err: %v", err.Error())))
		return
	}
}

func userValidation(user User) bool {
	if user.Email == "" || user.Password == "" || user.FirstName == "" || user.LastName == "" || user.Organization == "" || user.Country == "" || user.Status == "" || user.Category == "" {
		return false
	}
	return true
}

func generateSixDigitCode() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(900000) + 100000
}

func saveCodeToDB(email string, code int) error {
	query := `INSERT INTO otps (email, code, created_at) VALUES ($1, $2)`
	err := DB.Exec(query, email, code, time.Now()).Error
	return err
}

type User struct {
	Email        string `json:"email"`
	Password     string `json:"password"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Organization string `json:"organization"`
	Country      string `json:"country"`
	Status       string `json:"status"`
	Category     string `json:"category"`
}
