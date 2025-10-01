package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/harsh-karwar/udhaar-tracker-backend/internal/database"
	"github.com/harsh-karwar/udhaar-tracker-backend/internal/handlers"
	"github.com/harsh-karwar/udhaar-tracker-backend/internal/repository"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	database.ConnectDB()
	repository.InitRepository()
	repository.InitUserRepository() // <-- This was the missing line

	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	v1 := router.Group("/api/v1")
	{
		v1.POST("/signup", handlers.SignupHandler)
		v1.POST("/login", handlers.LoginHandler)
		v1.POST("/udhaars", handlers.CreateUdhaar)
		v1.GET("/udhaars", handlers.GetUdhaars)
		v1.GET("/udhaars/:id", handlers.GetUdhaarByID)
		v1.PUT("/udhaars/:id", handlers.UpdateUdhaar)
		v1.DELETE("/udhaars/:id", handlers.DeleteUdhaar)
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server listening on port %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}

