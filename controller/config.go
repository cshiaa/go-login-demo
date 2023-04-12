package controller

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"

	"github.com/cshiaa/go-login-demo/global"
	"github.com/cshiaa/go-login-demo/source/system"
)

var systemConfig = system.SystemConfig{}


func GetConfigList(c *gin.Context) {


	conf, err := systemConfig.GetSystemConfig()
	fmt.Println(global.RY_CONFIG)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message":"success","config": conf})
}

func UpadteConfigList(c *gin.Context) {

	var system = system.System{}

	global.RY_LOG.Info("前端请求更新的配置信息:")
	if err := c.ShouldBindJSON(&system); err!= nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
    }
	systemConfig.SetSystemConfig(system)
	c.JSON(http.StatusOK, gin.H{"message":"success"})
}