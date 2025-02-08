package main

import (
	"inventaris-api/db"
	"inventaris-api/routes"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
    // Koneksi ke database
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file")
    } 
    db.ConnectDatabase()

    // Setup router
    r := gin.Default()
    routes.SetupRoutes(r)

    // Jalankan server
    r.Run(":8111") // mendengarkan di port 8080
}
