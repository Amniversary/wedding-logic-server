package models

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/jinzhu/gorm"
	"fmt"
	"github.com/Amniversary/wedding-logic-server/config"
	"log"
)

var db *gorm.DB

func InitDataBase() {
	openDb()
}

func openDb() {
	db1, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&loc=Local",
		config.USER,
		config.PASS,
		config.HOST,
		config.DBName,
	))
	if err != nil {
		log.Printf("init DateBase error: [%v]", err)
		return
	}
	//db1.LogMode(true)
	db = db1
	//db.DB().SetMaxIdleConns(30)
	//db.DB().SetMaxOpenConns(100)
	initTable()
}

func initTable() {
	db.AutoMigrate(new(Card), new(Collection))
}
