package controller

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/cshiaa/go-login-demo/models"

	"github.com/cshiaa/go-login-demo/utils"
)


type LoginInput struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	
	var input LoginInput

	if err := c.ShouldBindJSON(&input); err!= nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
    }

	u := models.User{}

	u.Username = input.Username
	u.Password = input.Password

	atoken, rtoken, err := models.LoginCheck(u.Username, u.Password)
	if err!= nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username or password is incorrect"})
            return
	}

	c.JSON(http.StatusOK, gin.H{"atoken": atoken, "rtoken": rtoken})
}

func CurrentUser(c *gin.Context){

	user_id, err := utils.ExtractTokenID(c)
	
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	u,err := models.GetUserByID(user_id)
	
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message":"success","data":u})
}