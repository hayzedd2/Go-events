package routes

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/hayzedd2/Go-events/models"
	"github.com/hayzedd2/Go-events/utils"
)

func signUp(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse request data",
		})
		return
	}
	err = user.Save()
	if err != nil {
		switch {
		case strings.Contains(err.Error(), "UNIQUE constraint failed: users.email"):
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Email address is already registered",
			})
			return
		case strings.Contains(err.Error(), "UNIQUE constraint failed: users.userName"):
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Username is already taken",
			})
			return
		case strings.Contains(err.Error(), "password must be at least 8 characters long"):
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "password must be at least 8 characters long",
			})
			return
		default:
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "could not create account",
			})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "User created succesfully!",
	})

}

func login(c *gin.Context) {
	var user models.UserLogin
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse request data",
		})
		return
	}
	validatedUser, err := user.ValidateCredentials()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	token, err := utils.GenerateToken(user.Email, validatedUser.UserName, validatedUser.UserId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "login succesful", "token": token, "user": gin.H{
		"email":    user.Email,
		"userName": validatedUser.UserName,
		"userId":   validatedUser.UserId,
	}})
}
