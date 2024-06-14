package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func AddResearch(c *gin.Context) {
	var research Research
	err := c.ShouldBindJSON(&research)
	if err != nil {
		log.Println("Couldn't read data from request body, err = ", err)
		c.JSONP(400, "something went wrong. Please try again")
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

	err = c.SaveUploadedFile(file, file.Filename+" "+research.Category)
	if err != nil {
		log.Println("Couldn't save the file, because of err = ", err)
		c.JSONP(500, "Something went wrong")
		return
	}
	research.Location = file.Filename + " " + research.Category

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
