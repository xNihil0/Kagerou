package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var UserDB *gorm.DB

func init() {
	var err error
	dbLogger := logger.New(
		log.New(os.Stdout, "\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			Colorful:      false,
			LogLevel:      logger.Warn,
		})
	UserDB, err = gorm.Open(sqlite.Open("Kagerou.db"), &gorm.Config{
		Logger: dbLogger,
	})
	if err != nil {
		panic(err)
	}
	err = UserDB.AutoMigrate(&User{})
	if err != nil {
		panic(err)
	}
}
