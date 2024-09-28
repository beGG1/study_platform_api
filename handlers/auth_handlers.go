package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// TODO: Make it with hashing and jwt
func Login(c *gin.Context) {
	var authStuct AuthInput

	mockUser := "admin"
	mockPassword := "admin"

	err := c.ShouldBind(&authStuct)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if authStuct.Username != mockUser || authStuct.Password != mockPassword {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Wrong username or password"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": "Success"})
}
