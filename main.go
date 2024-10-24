package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hayzedd2/Go-events/db"
	"github.com/hayzedd2/Go-events/routes"
)

func main() {
	db.InitDB()
	gin.SetMode(gin.DebugMode)
	s := gin.Default()
	routes.RegisterRoutes(s)
	s.Run(":8080")

}
