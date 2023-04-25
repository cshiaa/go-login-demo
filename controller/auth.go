package controller

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/cshiaa/go-login-demo/models/system"
	"github.com/cshiaa/go-login-demo/common/response"

)

type RegisterInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}


func Register(c *gin.Context) {

	var input RegisterInput
	if err := c.ShouldBindJSON(&input); err!= nil {
		response.FailWithMessage("解析用户数据报错", c)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
	}

	u := system.User{}

	u.Username = input.Username
	u.Password = input.Password

	if err := u.BeforeSave(); err != nil {
		response.FailWithMessage("注册的用户信息有误", c)
        return
    }

	_, err := u.SaveUser()
	if err!= nil {
		response.FailWithMessage("用户保存失败", c)
        return
    }
	response.SuccessWithMessage("注册成功", c)
}