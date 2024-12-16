package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hayzedd2/eventease-be/db"
	"github.com/hayzedd2/eventease-be/routes"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func InitEnv() {
	err := godotenv.Load()
	if os.Getenv("ENVIRONMENT") == "DEVELOPMENT" {
		if err != nil {
			log.Println("Error loading .env file:", err)
		}
	}
}

func main() {
	InitEnv()
	db.InitDB()
	if os.Getenv("ENVIRONMENT") == "production" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}
	s := gin.Default()
	// s.Use(CORSMiddleware())
	routes.RegisterRoutes(s)
	s.Run(":8080")

}

func CORSMiddleware() gin.HandlerFunc {
	allowedOrigins := []string{
		"https://eventsease.vercel.app",
		"http://localhost:3000",
	}
	return func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")
		if contains(allowedOrigins, origin) {
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
		}
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}
