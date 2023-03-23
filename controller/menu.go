package controller

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/cshiaa/go-login-demo/models"

	"github.com/cshiaa/go-login-demo/utils"

)



func GetMenuList(c *gin.Context){

	userId, err := utils.ExtractTokenID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	menu, err := models.GetMenu(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message":"success","menuList": menu})
}
