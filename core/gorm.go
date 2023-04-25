package core

import (
	"os"
	"time"

	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"


	"go.uber.org/zap"

	gormlogger "gorm.io/gorm/logger"

	"github.com/cshiaa/go-login-demo/global"
	"github.com/cshiaa/go-login-demo/models/system"
	"github.com/cshiaa/go-login-demo/logger"

)

func InitDatabase() (*gorm.DB) {

	newLogger := gormlogger.New(
		// log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Writer{},
		gormlogger.Config{
		  SlowThreshold:              time.Second,   // Slow SQL threshold
		  LogLevel:                   4, // Log level
		  IgnoreRecordNotFoundError: true,           // Ignore ErrRecordNotFound error for logger
		  ParameterizedQueries:      false,           // Don't include params in the SQL log
		  Colorful:                  false,          // Disable color
		},
	)

	m := global.RY_CONFIG.Mysql
	if m.Dbname == "" {
		return nil
	}
	DBURL := m.Dsn()
	db, err := gorm.Open(mysql.Open(DBURL), &gorm.Config{
		Logger:newLogger,
	})
	if err!= nil {
		global.RY_LOG.Sugar().Fatalf("Error opening database: ", err)
	}

	return db
}

// RegisterTables 注册数据库表专用
// Author SliverHorn
func RegisterTables() {

	db := global.RY_DB

	err := db.AutoMigrate(&system.User{}, &system.Menu{}, &system.RolePermissions{})

	if err != nil {
		global.RY_LOG.Error("register table failed", zap.Error(err))
		os.Exit(0)
	}
	global.RY_LOG.Info("register table success")
}
