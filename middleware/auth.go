package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hayzedd2/Go-events/utils"
)

func Authenticate(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	if token == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "could not authenticate user"})
		return
	}
	userId, err := utils.VerifyToken(token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "could not verify token"})
		return
	}
	c.Set("userId", userId)
	c.Next()
}
