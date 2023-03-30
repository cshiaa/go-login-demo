package core

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"


	"go.uber.org/zap"

	// "moul.io/zapgorm2"

	zapgorm2 "github.com/cshiaa/go-login-demo/logger"
	"github.com/cshiaa/go-login-demo/global"
	"github.com/cshiaa/go-login-demo/models"

)

func InitDatabase() (*gorm.DB) {

	loggorm := zapgorm2.New(zap.L())
	loggorm.SetAsDefault() // optional: configure gorm to use this zapgorm.Logger for callbacks
	
	m := global.RY_CONFIG.Mysql
	if m.Dbname == "" {
		return nil
	}
	DBURL := m.Dsn()
	db, err := gorm.Open(mysql.Open(DBURL), &gorm.Config{
		Logger:loggorm,
		//Logger:logger.Default.LogMode(logger.Info),
	})
	if err!= nil {
		fmt.Println("Error opening database: %v", m.Dirver)
		log.Fatalf("Error opening database: ", err)
	} else {
		fmt.Println("Connected to database: ", m.Dirver)
	}

	return db
}

// RegisterTables 注册数据库表专用
// Author SliverHorn
func RegisterTables() {
	db := global.RY_DB

	err := db.AutoMigrate(&models.User{}, &models.Menu{}, &models.Role_Permissions{})

	if err != nil {
		global.RY_LOG.Error("register table failed", zap.Error(err))
		os.Exit(0)
	}
	global.RY_LOG.Info("register table success")
}
