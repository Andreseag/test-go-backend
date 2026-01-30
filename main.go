package main

import (
	"net/http"
	"time"

	"github.com/Andreseag/test-go-backend/config"
	"github.com/Andreseag/test-go-backend/controllers"
	"github.com/Andreseag/test-go-backend/models"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Middleware para habilitar CORS
func enableCORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Permitir que React (o cualquier origen) acceda
		w.Header().Set("Access-Control-Allow-Origin", "*") 
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Manejar la petición "pre-flight" (el navegador pregunta antes de enviar datos)
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next(w, r)
	}
}

func main() {
	config.ConectarDB()
	config.DB.AutoMigrate(&models.Task{})

	r := gin.Default()

	// Configuración de CORS profesional
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		MaxAge:           12 * time.Hour,
	}))
 
	// Envolvemos las rutas con el middleware
  r.GET("/api/tasks", controllers.GetTasks)
	r.POST("/api/tasks/new", controllers.CreateTask)
	
	// ¡Aquí está lo que pediste! El :id es el parámetro
	r.PUT("/api/tasks/:id", controllers.UpdateTask)
	r.DELETE("/api/tasks/:id", controllers.DeleteTask) 

	r.Run(":8080")
}