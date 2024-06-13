package main

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {
	var user UserLogin
	if err := c.ShouldBindJSON(&user); err != nil {
		log.Println("failed to get data from request, err = ", err)
		c.JSONP(400, "please try again")
		return
	}
	if err := userExists(user.Email, user.Password); err != nil {
		log.Println("failed to get data from db, err = ", err)
		c.JSONP(400, "please try again")
		return
	}
	token, err := getToken(user.Email)
	if err != nil {
		log.Println("failed to initialize token, err = ", err)
		c.JSONP(400, "please try again")
		return
	}
	if err = tokenToDb(token); err != nil {
		log.Println("failed to save token to db")
		c.JSONP(500, "sorry, something went wrong. please try again")
	}

	c.JSONP(200, token)
}

func userExists(email, password string) error {
	var user UserDAO
	query := `select * from users where email = $1 and password = $2 and activated = true`
	tx := db.Raw(query, email, password).Scan(&user)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("nothing found")
	}
	return nil
}

func getToken(email string) (Token, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	calms := token.Claims.(jwt.MapClaims)
	calms["email"] = email
	calms["time"] = time.Now()
	secretKey := []byte("secret")
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return Token{}, err
	}
	var NewToken Token
	NewToken.Token = tokenString
	NewToken.Email = email
	NewToken.ExpirationTime = time.Now().Add(1 * time.Hour)
	return NewToken, nil
}

func tokenToDb(token Token) error {
	query := `insert into tokens (email, token) values ($1, $2)`
	err := db.Exec(query, token.Email, token.Token).Error
	return err
}

type UserDAO struct {
	UserID       int    `json:"user_id"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Organization string `json:"organization"`
	Country      string `json:"country"`
	Status       string `json:"status"`
	Category     string `json:"category"`
	activated    string `json:"activated"`
}

type Token struct {
	ID             int       `json:"id"`
	Token          string    `json:"token"`
	ExpirationTime time.Time `json:"expiration_time"`
	Email          string    `json:"user_id"`
}
