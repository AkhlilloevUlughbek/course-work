package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path/filepath"
)

func DownloadResearch(c *gin.Context) {
	category := c.GetHeader("category")
	title := c.GetHeader("title")

	// Проверка наличия заголовков
	if category == "" || title == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Category and title headers are required"})
		return
	}

	// Формируем путь к файлу
	filename := title + " " + category
	filepath := filepath.Join(filename)

	// Проверяем, существует ли файл
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		c.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
		return
	}

	// Устанавливаем заголовок Content-Disposition для скачивания файла
	c.Header("Content-Disposition", "attachment; filename="+filename)
	c.File(filepath)
}
