package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/hayzedd2/Go-events/utils"
)


func Authenticate(c *gin.Context) {
	authHeader := c.Request.Header.Get("Authorization")
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "could not authenticate user"})
		return
	}
	token := strings.TrimPrefix(authHeader, "Bearer ")
	userId, err := utils.VerifyToken(token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Could not verify token"})
		return
	}
	c.Set("userId", userId)
	c.Next()
}
