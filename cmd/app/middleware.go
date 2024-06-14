package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, PATCH, GET, PUT, DELETE")
		c.Next()
	}
}

func CheckUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Token")
		email := c.GetHeader("Email")
		err := getTokenFromDB(token, email)
		if err != nil {
			log.Println("Token not found, err = ", err)
			c.JSONP(http.StatusUnauthorized, "Sorry you need to enter you account one more time")
			return
		}
		c.Next()
	}

}

func getTokenFromDB(email, token string) error {
	var tok Token
	query := `select * from tokens where email=$1 and token=$2`
	tx := db.Raw(query, email, token).Scan(&tok)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("nothing here")
	}

	if time.Now().Sub(tok.CreatedAt) < 24*time.Hour {
		return errors.New("token invalid")
	}

	return nil
}
