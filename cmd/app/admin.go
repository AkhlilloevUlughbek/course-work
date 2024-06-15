package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func getNotApproved(c *gin.Context) {
	researches, err := getNotApprFromDb()
	if err != nil {
		log.Println("Something went wrong while tried to get not approved researches from db, error = ", err)
		c.JSON(400, gin.H{"error": err})
		return
	}

	c.JSON(200, researches)
}

func getNotApprFromDb() ([]Research, error) {
	query := `select * from researches where approved=false`
	var researches []Research
	err := db.Raw(query).Scan(&researches).Error
	if err != nil {
		return nil, err
	}

	return researches, nil
}

func approve(c *gin.Context) {
	title := c.GetHeader("title")
	query := `update researches set approved=true where title=$1`
	err := db.Exec(query, title).Error
	if err != nil {
		log.Println("faced some problems while tried to approve, error = ", err)
		c.JSON(400, gin.H{"error": err})
		return
	}

	c.JSON(200, "ok")
}
