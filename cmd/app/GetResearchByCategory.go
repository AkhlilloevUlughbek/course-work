package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func getResearchesByCategories(c *gin.Context) {
	category := c.GetHeader("category")
	researches, err := getResearchesFromDb(category)
	if err != nil {
		log.Println("couldn't get researches from db by category, error = ", err)
		c.JSON(400, gin.H{
			"error": "wrong category sent",
		})
	}
	c.JSON(200, researches)
}

func getResearchesFromDb(category string) ([]Research, error) {
	query := `select * from researches where category = $1 and approved=true`
	var researches []Research
	err := db.Raw(query, category).Scan(&researches).Error
	if err != nil {
		return nil, err
	}
	return researches, nil
}
