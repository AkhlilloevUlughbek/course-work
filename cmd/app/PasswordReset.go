package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/gomail.v2"
	"log"
)

type ResetUser struct {
	Password string `json:"password"`
}

func ResetPassword(c *gin.Context) {
	email := c.Param("email")
	if email == "" {
		log.Println("empty email")
		c.JSONP(400, "empty email")
		return
	}

	code := generateSixDigitCode()
	if err := saveCodeToDB(email, code); err != nil {
		log.Printf("failed to save otp to db, err = %v", err)
		c.JSONP(500, "something went wrong. please try again")
		return
	}

	if err := sendToUserForReset(email, code); err != nil {
		if err = deleteCodeFromDB(email); err != nil {
			log.Printf("failed to delete record from db: %v", err)
			c.JSON(500, "something went wrong. Please reach support")
			return
		}
		log.Printf("failed send otp to user: %v", err)
		c.JSON(500, "something went wrong. Please try again")
		return
	}
	c.JSONP(200, "Everything is ok)")
}

func FinishReset(c *gin.Context) {
	email := c.Param("email")
	if email == "" {
		log.Println("empty email")
		c.JSONP(400, "empty email")
		return
	}
	var res ResetUser
	if err := c.ShouldBindJSON(&res); err != nil {
		log.Printf("failed get data for reset, err: %v", err)
		c.JSON(400, "something went wrong. Please  enter email, and new pass again")
		return
	}

	if err := updateUser(email, res.Password); err != nil {
		log.Println("failed to save data to db, err = ", err)
		c.JSONP(400, "something went wrong. please enter email, and new pass again")
		return
	}

	c.JSONP(200, "Everything is okay")
}

func updateUser(email, password string) error {
	query := `update users set password=$1 where email = $2;`
	err := db.Exec(query, password, email).Error
	return err
}

func sendToUserForReset(email string, otp int) error {
	m := gomail.NewMessage()
	m.SetHeader("From", "lutfullomelikov686@gmail.com")
	m.SetHeader("To", email)
	m.SetHeader("Subject", "Подтверждение регистрации")
	m.SetBody("text/html", fmt.Sprintf(`
  <p>Здраствуйте</p>
  <p>Для сброса пароля введите этот код: %d</p>
  <p>Никому не передавайте код.</p>
`, otp))
	d := gomail.NewDialer("smtp.gmail.com", 587, "lutfullomelikov686@gmail.com", "otsz ytlc zsdx rzdu")

	if err := d.DialAndSend(m); err != nil {
		log.Println("ERORRRRR: %v", err)
		return err
	}

	return nil
}
