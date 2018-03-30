package mysql

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/jinzhu/gorm"
	"github.com/Amniversary/wedding-logic-server/config"
	"fmt"
	"log"
)

var db *gorm.DB

func NewMysql(c *config.Config) {
	openDb(c)
}

func openDb(c *config.Config) {
	db1, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&loc=Local",
		c.DBInfo.User,
		c.DBInfo.Pass,
		c.DBInfo.Host,
		c.DBInfo.DBName,
	))
	if err != nil {
		log.Printf("init DateBase error: [%v]", err)
		return
	}
	if c.DBDebug {
		db1.LogMode(true)
	}

	db = db1
	db.DB().SetMaxIdleConns(20)
	db.DB().SetMaxOpenConns(50)
	initTable()
}

func initTable() {
	db.AutoMigrate(
		new(Card),
		new(Collection),
		new(Production),
		new(ClickProduction),
		new(SmsMessage),
		new(Schedule),
		new(Cooperation),
		)
}
