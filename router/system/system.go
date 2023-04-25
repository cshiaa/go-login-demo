package system

import (
	
	"github.com/gin-gonic/gin"
	"github.com/cshiaa/go-login-demo/controller"
	"github.com/cshiaa/go-login-demo/middlewares"
)

func Routers(e *gin.Engine) {

	e.Use(middlewares.JwtAuthMiddleware())

	menu := e.Group("/menu")
	{
		menu.GET("/allMenu", controller.GetMenu)
		menu.POST("/asyncMenu", controller.GetMenuList)
		menu.POST("/getUserMenu", controller.GetUserMenuList)
		menu.POST("/updateUserMenu", controller.UpdateUserMenu)
	}

	user := e.Group("/user")
	{
		user.GET("/list", controller.GetUserList)
	}

	config := e.Group("/config")
	{
		config.GET("/list", controller.GetConfigList)
		config.POST("/update", controller.UpadteConfigList)
	}

	file := e.Group("/file")
	{
		file.POST("/upload", controller.UploadFile)
		file.GET("/getConfigFile", controller.GetFile)
	}

	kubernetes := e.Group("/kubernetes")
	{
		kubernetes.GET("/version", controller.GetKubernetesVersion)
		kubernetes.GET("/resources/get", controller.GetKubernetesResource)

	}
}