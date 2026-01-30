package main

import (
	"log"
	"net/http"

	"github.com/Andreseag/test-go-backend/config"
	"github.com/Andreseag/test-go-backend/controllers"
	"github.com/Andreseag/test-go-backend/models"
)

// Middleware para habilitar CORS
func enableCORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Permitir cualquier origen (en producción pondrías tu URL de React)
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Si es una petición de tipo OPTIONS (pre-flight), respondemos OK y salimos
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next(w, r)
	}
}

func main() {
	// Conectamos a la DB
	config.ConectarDB()
  config.DB.AutoMigrate(&models.Producto{})

	// Definimos las rutas
	http.HandleFunc("/api/productos", enableCORS(controllers.GetProductos)) // GET para leer
	http.HandleFunc("/api/productos/nuevo", enableCORS(controllers.CrearProducto)) // POST para crear

	log.Println("🚀 Servidor corriendo en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}