package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func PersAccount(c *gin.Context) {
	email, exists := c.Get("email")
	if !exists {
		log.Println("email is not found in context")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email not found"})
		return
	}
	emailStr, exists := email.(string)
	if !exists {
		log.Println("Couldn't transform email to string")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email not found"})
		return
	}

	user, err := getUserFromDB(emailStr)
	if err != nil {
		log.Println("couldn't get user from db, err = ", err)
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}

	c.JSON(200, user)
}

func getUserFromDB(email string) (User, error) {
	var user User
	query := `select * from users where email = $1`
	err := db.Raw(query, email).Scan(&user).Error
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func getResearches(c *gin.Context) {
	email, exists := c.Get("email")
	if !exists {
		log.Println("email is not found in context")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email not found"})
		return
	}

	emailStr, exists := email.(string)
	if !exists {
		log.Println("Couldn't transform email to string")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email not found"})
		return
	}

	researches, err := getResearchFromDB(emailStr)
	if err != nil {
		log.Println("Couldn't get researches from db, err = ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email not found"})
		return
	}

	c.JSON(200, researches)
}

func getResearchFromDB(email string) ([]Research, error) {
	var researches []Research
	query := `select * from researches where user_email = $1 and approved=true`
	err := db.Raw(query, email).Scan(&researches).Error
	if err != nil {
		return nil, err
	}
	return researches, nil
}
