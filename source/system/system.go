package system

import (
	"github.com/cshiaa/go-login-demo/global"
	"github.com/cshiaa/go-login-demo/source/config"
	"github.com/cshiaa/go-login-demo/utils"
)

// 配置文件结构体
type System struct {
	Config config.Server `json:"config"`
}

// 配置文件结构体
type SystemConfig struct {
	Config config.Server `json:"config"`
}

func (systemConfig *SystemConfig) GetSystemConfig() (conf config.Server, err error) {
	return global.RY_CONFIG, nil
}

// @description   set system config,
//@author: [piexlmax](https://github.com/piexlmax)
//@function: SetSystemConfig
//@description: 设置配置文件
//@param: system model.System
//@return: err error

func (systemConfig *SystemConfig) SetSystemConfig(system System) (err error) {
	cs := utils.StructToMap(system.Config)
	for k, v := range cs {
		global.RY_VP.Set(k, v)
	}
	err = global.RY_VP.WriteConfig()
	return err
}