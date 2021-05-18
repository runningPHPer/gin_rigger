package rigger

import (
	"log"
	"xorm.io/xorm"
)

type XormAdapter struct {
	*xorm.Engine
}

func NewXormAdapter() *XormAdapter {
	engine, err := xorm.NewEngine("mysql", "root:123456@tcp(127.0.0.1:3306)/oAuth2?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
	}
	engine.SetMaxIdleConns(5)
	engine.SetMaxOpenConns(10) //最大连接数
	return &XormAdapter{Engine: engine}
}
