package models

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDataBase(){
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading.env file: %v", err)
	}

	db_driver := os.Getenv("DBDIRVER")
	db_host := os.Getenv("DBHOST")
	db_user := os.Getenv("DBUSER")
	db_password := os.Getenv("DBPASSWORD")
	db_name := os.Getenv("DBNAME")
	db_port := os.Getenv("DBPORT")

	DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", db_user, db_password, db_host, db_port, db_name)

	DB, err = gorm.Open(mysql.Open(DBURL), &gorm.Config{
		Logger:logger.Default.LogMode(logger.Info),
	})
	if err!= nil {
		fmt.Println("Error opening database: %v", db_driver)
		log.Fatalf("Error opening database: ", err)
	} else {
		fmt.Println("Connected to database: ", db_driver)
	}

	DB.AutoMigrate(&User{}, &Menu{}, &Role_Permissions{})
}