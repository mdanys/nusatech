package databases

import (
	"fmt"
	"nusatech/config"
	"nusatech/features/users/repository"

	"github.com/labstack/gommon/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(c *config.AppConfig) *gorm.DB {
	str := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		c.DBUser,
		c.DBPass,
		c.DBHost,
		c.DBPort,
		c.DBName,
	)

	db, err := gorm.Open(mysql.Open(str), &gorm.Config{})
	if err != nil {
		log.Error("db config error: ", err.Error())
		return nil
	}

	return db
}

func MigrateDB(db *gorm.DB) {
	db.AutoMigrate(&repository.User{})
	db.AutoMigrate(&repository.Balance{})
	db.AutoMigrate(&repository.Currency{})
	db.AutoMigrate(&repository.Mailer{})
}
