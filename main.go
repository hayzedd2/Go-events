package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/hayzedd2/Go-events/db"
	"github.com/hayzedd2/Go-events/routes"
)

func main() {
	db.InitDB()
	if os.Getenv("ENVIRONMENT") == "production" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	s := gin.Default()
	s.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})
	routes.RegisterRoutes(s)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	s.Run(":" + port)

}
