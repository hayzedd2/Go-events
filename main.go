package main

import (
	"log"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/hayzedd2/Go-events/db"
	"github.com/hayzedd2/Go-events/routes"
	"github.com/joho/godotenv"
)

func InitEnv() {
	if os.Getenv("ENVIRONMENT") == "development" {
        err := godotenv.Load()
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
	s.Use(CORSMiddleware())
	routes.RegisterRoutes(s)
	s.Run(":8080")

}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

func CORSMiddleware() gin.HandlerFunc {
	allowedOrigins := strings.Split(os.Getenv("ALLOWED_ORIGIN"), ",")
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
