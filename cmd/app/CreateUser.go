package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/gomail.v2"
	"log"
	"math/rand"
	"time"
)

func CreateUser(c *gin.Context) {
	var newUser User
	err := c.ShouldBindJSON(&newUser)
	if err != nil {
		log.Printf("failed to bind json to user: %v", err)
		c.JSON(400, "something went wrong. Please try again")
		return
	}

	if !userValidation(newUser) {
		c.JSON(400, "you forgot to sent some relevant field(s). Please, check everything and try again")
	}

	otp := generateSixDigitCode()

	if err = saveCodeToDB(newUser.Email, otp); err != nil {
		log.Printf("failed to save user: %v", err)
		c.JSON(500, "something went wrong. Please try again later")
		return
	}

	if err = sendToUser(newUser.Email, otp); err != nil {
		if err = deleteCodeFromDB(newUser.Email); err != nil {
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
	query := `INSERT INTO otps (email, otp) VALUES ($1, $2)`
	err := db.Exec(query, email, code).Error
	return err
}

func deleteCodeFromDB(email string) error {
	query := `delete from otps where email = $1`
	err := db.Exec(query, email).Error
	return err
}

func sendToUser(email string, otp int) error {
	m := gomail.NewMessage()
	m.SetHeader("From", "lutfullomelikov686@gmail.com")
	m.SetHeader("To", email)
	m.SetHeader("Subject", "Подтверждение регистрации")
	m.SetBody("text/html", fmt.Sprintf(`
  <p>Здраствуйте</p>
  <p>Для подтверждаения регистрации введите этот код: %d</p>
  <p>Никому не передавайте код.</p>
`, otp))
	d := gomail.NewDialer("smtp.gmail.com", 587, "lutfullomelikov686@gmail.com", "otsz ytlc zsdx rzdu")

	if err := d.DialAndSend(m); err != nil {
		log.Println("ERORRRRR: %v", err)
		return err
	}

	return nil
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
