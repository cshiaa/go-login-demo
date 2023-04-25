package core

import (
	"github.com/gin-gonic/gin"
	"github.com/cshiaa/go-login-demo/middlewares"
	"github.com/cshiaa/go-login-demo/logger"
	"github.com/cshiaa/go-login-demo/router/api"
	"github.com/cshiaa/go-login-demo/router/system"
	"github.com/cshiaa/go-login-demo/global"
)

type Option func(*gin.Engine)

var options = []Option{}

// 注册app的路由配置
func include(opts ...Option) {
	options = append(options, opts...)
}

// 初始化
func ginInit() *gin.Engine {
	router := gin.Default()
	router.Use(middlewares.Cors())
	router.Use(logger.GinLogger(), logger.GinRecovery(true))
	for _, opt := range options {
		opt(router)
	}
	return router
}

func RunServer() {
	// 加载多个APP的路由配置
	include(api.Routers, system.Routers)
	// 初始化路由
	r := ginInit()
	if err := r.Run(":8089"); err != nil {
		global.RY_LOG.Sugar().Fatalf("startup service failed, err:%v\n", err)
	}
}