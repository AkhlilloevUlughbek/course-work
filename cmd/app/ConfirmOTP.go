package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

type Email struct {
	Email string `json:"email"`
}

func ConfirmOTP(c *gin.Context) {
	otp := c.Param("otp")
	if otp == "" {
		log.Println("empty otp")
		c.JSONP(400, "empty otp")
		return
	}
	var em Email
	err := c.ShouldBindJSON(&em)
	if err != nil {
		log.Println("empty email")
		c.JSONP(400, "empty email")
		return
	}

	fmt.Println(em)
	awaitedOTP, err := getOTPFromDb(em.Email)
	if err != nil {
		log.Println("Failed to retrieve otp info from db, err = ", err)
		c.JSONP(500, "something went wrong, please try again later")
		return
	}
	duration := time.Now().Sub(awaitedOTP.CreatedAt)
	if duration.Minutes() > 5 {
		log.Println("OTP valid time expired")
		if err = deleteCodeFromDB(em.Email); err != nil {
			log.Printf("failed to delete record from db: %v", err)
			c.JSON(500, "something went wrong. Please reach support")
			return
		}
		if err = deleteUserFromDB(em.Email); err != nil {
			log.Printf("failed to delete record from db: %v", err)
			c.JSON(500, "something went wrong. Please reach support")
			return
		}
		c.JSONP(400, "otp expired. Try again registration")
		return
	}

	if otp == fmt.Sprint(awaitedOTP.OTP) {
		if err = updateUsers(em.Email); err != nil {
			log.Println("Failed to update the column activated for users, err = ", err)
			c.JSONP(500, "something went wrong. Try again")
			return
		}
		if err = deleteCodeFromDB(em.Email); err != nil {
			log.Println("failed to delete otp from otps, err = ", err)
			c.JSONP(200, "Everything is great")
			return
		}
		c.JSONP(200, "Everything is great")
		return
	}
	log.Printf("invalid otp got = %v want = %v", otp, awaitedOTP.OTP)
	c.JSONP(400, "Invalid otp. Try again")
	return
}

func getOTPFromDb(email string) (OTP, error) {
	var otp OTP
	query := `select * from otps where email = $1;`
	err := db.Raw(query, email).Scan(&otp).Error
	return otp, err
}

func updateUsers(email string) error {
	query := `update users set activated=true where email = $1;`
	err := db.Exec(query, email).Error
	return err
}

type OTP struct {
	Email     string
	OTP       int
	CreatedAt time.Time
}
