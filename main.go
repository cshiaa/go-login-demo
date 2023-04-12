package main


import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/cshiaa/go-login-demo/core"
	"github.com/cshiaa/go-login-demo/global"

	"github.com/cshiaa/go-login-demo/controller"
	"github.com/cshiaa/go-login-demo/middlewares"
	"github.com/cshiaa/go-login-demo/logger"

	// "time"
	// ginzap "github.com/gin-contrib/zap"
	// "go.uber.org/zap"



)

func main() {

	//配置文件初始化设置
	global.RY_VP = core.ViperInit()

	// init logger
	if err := logger.InitLogger(); err != nil {
		fmt.Printf("init logger failed, err:%v\n", err)
		return
	}

	//init mysql database
	global.RY_DB = core.InitDatabase()
	if global.RY_DB != nil {
		core.RegisterTables() // 初始化表
		// 程序结束前关闭数据库链接
		db, _ := global.RY_DB.DB()
		defer db.Close()
	}

	router := gin.Default()
	router.Use(middlewares.Cors())
	router.Use(logger.GinLogger(), logger.GinRecovery(true))

	public := router.Group("/api")
	public.POST("/register", controller.Register)
	public.POST("/login", controller.Login)

	protected := router.Group("/api/admin")
	protected.Use(middlewares.JwtAuthMiddleware())
	protected.GET("/user", controller.CurrentUser)
	protected.GET("/menu", controller.GetMenuList)
	protected.POST("/asyncMenu", controller.GetMenuList)

	protectedUser := router.Group("/api/user")
	{
		protectedUser.GET("/list", controller.GetUserList)
		protectedUser.GET("/add", controller.CurrentUser)
	}
	protectedUser.Use(middlewares.JwtAuthMiddleware())

	protectedConfig := router.Group("/api/config")
	{
		protectedConfig.GET("/list", controller.GetConfigList)
		protectedConfig.POST("/update", controller.UpadteConfigList)
	}
	protectedConfig.Use(middlewares.JwtAuthMiddleware())

	protectedFile := router.Group("/file")
	{
		protectedFile.POST("/upload", controller.UploadFile)
		protectedFile.GET("/getFile", controller.GetFile)

	}
	protectedFile.Use(middlewares.JwtAuthMiddleware())

	protectedKubernetes := router.Group("/kubernetes")
	{
		protectedKubernetes.GET("/version", controller.GetKubernetesVersion)
		protectedKubernetes.GET("/resources/get", controller.GetKubernetesResource)

	}
	protectedKubernetes.Use(middlewares.JwtAuthMiddleware())
	// Listen and Server in 0.0.0.0:8080
	router.Run(":8089")
}