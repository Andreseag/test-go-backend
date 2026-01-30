package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Andreseag/test-go-backend/config"
	"github.com/Andreseag/test-go-backend/models"
)

func GetProductos(w http.ResponseWriter, r *http.Request) {
	var productos []models.Producto
	// Buscamos todos los productos en la DB
	config.DB.Find(&productos)
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(productos)
}

func CrearProducto(w http.ResponseWriter, r *http.Request) {
	var p models.Producto
	// Decodificamos lo que viene del frontend (React)
	json.NewDecoder(r.Body).Decode(&p)
	
	// Guardamos en la DB
	config.DB.Create(&p)
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(p)
}