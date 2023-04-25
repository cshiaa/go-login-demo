package main


import (
	"fmt"

	"github.com/cshiaa/go-login-demo/core"
	"github.com/cshiaa/go-login-demo/global"
	"github.com/cshiaa/go-login-demo/logger"

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

	core.RunServer()
}