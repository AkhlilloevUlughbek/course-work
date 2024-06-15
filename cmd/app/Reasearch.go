package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func AddResearch(c *gin.Context) {
	var research Research
	jsonData := c.PostForm("json")
	err := json.Unmarshal([]byte(jsonData), &research)
	if err != nil {
		log.Println("couldn't parse json, err = ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}
	if !researchValidate(research) {
		log.Println("Forgot to add something crucial")
		c.JSONP(400, "Forgot to add something crucial")
		return
	}

	file, err := c.FormFile("uploadFile")
	if err != nil {
		log.Println("Couldn't read the file, because of err = ", err)
		c.JSONP(400, "Something went wrong")
		return
	}

	err = c.SaveUploadedFile(file, research.Title+" "+research.Category)
	if err != nil {
		log.Println("Couldn't save the file, because of err = ", err)
		c.JSONP(500, "Something went wrong")
		return
	}
	research.Location = research.Title + " " + research.Category

	err = saveResearchToDB(research)
	if err != nil {
		log.Println("Something went wrong because of ", err)
		c.JSONP(400, "Something went wring")
		return
	}

	// Файл успешно сохранен
	c.JSONP(200, "File uploaded successfully")
}

func researchValidate(research Research) bool {
	if research.Title == "" || research.Category == "" || research.Description == "" || research.UserEmail == "" {
		return false
	}
	return true
}

func saveResearchToDB(research Research) error {
	query := `insert into researches (title, description, location, user_email, category) values ($1, $2, $3, $4, $5)`
	err := db.Exec(query, research.Title, research.Description, research.Location, research.UserEmail, research.Category).Error
	if err != nil {
		return err
	}
	return err
}

type Research struct {
	ResearchID  int    `json:"research_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Location    string `json:"location"`
	UserEmail   string `json:"user_email"`
	Category    string `json:"category"`
	Approved    bool   `json:"approved"`
}
