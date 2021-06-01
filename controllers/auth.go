package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/Bastorx/test_lbc/models"
)

func Login(c *gin.Context) {
	c.Status(http.StatusNoContent)
}

func Logout(c *gin.Context) {
	c.Status(http.StatusNoContent)
}

func SignIn(c *gin.Context) {
	var form models.User

	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
}