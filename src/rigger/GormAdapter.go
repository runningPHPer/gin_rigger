package rigger

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

type GormAdapter struct {
	*gorm.DB
}

//数据库连接
func NewGormAdapter() *GormAdapter {
	dsn := "gorm:gorm@tcp(127.0.0.1:3306)/gorm?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: dsn, // DSN data source name
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal(err.Error())
	}
	return &GormAdapter{db}
}
