package controller

import (
	"net/http"
	"github.com/gin-gonic/gin"
	systemModel "github.com/cshiaa/go-login-demo/models/system"
	"github.com/cshiaa/go-login-demo/source/system"

	"github.com/cshiaa/go-login-demo/utils"
	"github.com/cshiaa/go-login-demo/common/response"
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

	u := systemModel.User{}

	u.Username = input.Username
	u.Password = input.Password

	atoken, rtoken, err := system.LoginCheck(u.Username, u.Password)
	if err!= nil {
		response.FailWithMessage("密码错误", c)
		// c.JSON(http.StatusBadRequest, gin.H{"error": "Username or password is incorrect"})
        return
	}
	token := make(map[string]interface{})
	token["atoken"] = atoken
	token["rtoken"] = rtoken
	response.SuccessWithDetailed(token, "登录成功", c)
	// c.JSON(http.StatusOK, gin.H{"atoken": atoken, "rtoken": rtoken})
}

func CurrentUser(c *gin.Context){

	user_id, err := utils.ExtractTokenID(c)
	
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	u,err := system.GetUserByID(user_id)
	
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message":"success","data":u})
	
}