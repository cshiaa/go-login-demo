package core

import (
	"fmt"

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
		fmt.Println("config file changed:", e.Name)
		if err = v.Unmarshal(&global.RY_CONFIG); err != nil {
			fmt.Println(err)
		}
	})
	if err = v.Unmarshal(&global.RY_CONFIG); err != nil {
		fmt.Println(err)
	}
	return v
}