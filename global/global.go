package global

import (

	"go.uber.org/zap"

	"github.com/spf13/viper"

	"gorm.io/gorm"

	"github.com/cshiaa/go-login-demo/source/config"
)

var (
	RY_DB     *gorm.DB
	RY_DBList map[string]*gorm.DB
	RY_VP     *viper.Viper
	RY_LOG    *zap.Logger
	RY_CONFIG *config.Server
)