package models

import (
	_ "github.com/go-sql-driver/mysql"
	"log"
	"xorm.io/xorm"
)

var Engine = Init()

func Init() *xorm.Engine {
	engine, err := xorm.NewEngine("mysql", "root:1234@tcp(127.0.0.1:3306)/cloud-disk?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Printf("Xorm New Engine Error", err)
		return nil
	}

	return engine

}
