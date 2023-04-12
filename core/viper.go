package core

import (
	"fmt"
	
	"go.uber.org/zap"

	"github.com/spf13/viper"
	"github.com/fsnotify/fsnotify"
	"github.com/cshiaa/go-login-demo/global"

)

func ViperInit() (*viper.Viper) {

	v := viper.New()
	var config = "./config.yaml"
	v.SetConfigFile(config)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		global.RY_LOG.Info("配置文件更改", zap.String("configName", e.Name))
		fmt.Println("config file changed:", e.Name)
		if err = v.Unmarshal(&global.RY_CONFIG); err != nil {
			global.RY_LOG.Error(err.Error())
		}
	})
	if err = v.Unmarshal(&global.RY_CONFIG); err != nil {
		global.RY_LOG.Error(err.Error())
	}
	return v
}