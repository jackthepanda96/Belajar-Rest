package mysql

import (
	"fmt"
	"log"

	"github.com/jackthepanda96/Belajar-Rest.git/config"
	"github.com/jackthepanda96/Belajar-Rest.git/feature/book/data"
	userData "github.com/jackthepanda96/Belajar-Rest.git/feature/user/data"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(cfg *config.AppConfig) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True", cfg.Username, cfg.Password, cfg.Address, cfg.Port, cfg.Name)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Cannot connect to DB")
	}

	return db
}

func MigrateData(db *gorm.DB) {
	db.AutoMigrate(userData.User{}, data.Book{})
}
