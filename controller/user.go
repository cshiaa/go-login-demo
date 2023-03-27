package controller

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/cshiaa/go-login-demo/models"
)

func GetUserList(c *gin.Context){


	users, err := models.GetAllUser()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message":"success","userList": users})
}