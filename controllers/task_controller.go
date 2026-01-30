package controllers

import (
	"net/http"

	"github.com/Andreseag/test-go-backend/config"
	"github.com/Andreseag/test-go-backend/models"
	"github.com/gin-gonic/gin"
)

func GetTasks(c *gin.Context) {
	var tasks []models.Task
	config.DB.Order("created_at desc").Find(&tasks)
	
	// Gin se encarga de los headers y de convertir a JSON
	c.JSON(http.StatusOK, tasks)
}

func CreateTask(c *gin.Context) {
	var t models.Task
	
	// ShouldBindJSON es el equivalente a Decode
	if err := c.ShouldBindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	// VALIDACIÓN
	if t.Status != models.StatusTodo && 
	   t.Status != models.StatusInProgress && 
	   t.Status != models.StatusDone {
		if t.Status == "" {
			t.Status = models.StatusTodo
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Estado no válido"})
			return
		}
	}
	
	config.DB.Create(&t)
	c.JSON(http.StatusCreated, t)
}

func UpdateTask(c *gin.Context) {
	id := c.Param("id") 
	
	var task models.Task
	if err := config.DB.First(&task, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Tarea no encontrada"})
		return
	}

	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	config.DB.Save(&task)
	c.JSON(http.StatusOK, task)
}