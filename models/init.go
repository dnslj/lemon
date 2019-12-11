package models

import (
	"fmt"
	"log"
	"os"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"lemon/utils/logging"
)

type Database struct {
	Local *gorm.DB
}

var DB *Database

// 打开数据库链接
// https://www.jianshu.com/p/f7419395e4cc
func openDB(username, password, addr, name string) *gorm.DB {
	config := fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=%s",
		username,
		password,
		addr,
		name,
		"Local")

	db, err := gorm.Open("mysql", config)

	if err != nil {
		logging.Error(err, "Database connection failed. Database name: %s", name)
	}

	// Gorm有内置的日志记录器支持，默认情况下，它会打印发生的错误
	db.LogMode(true)

	filePath := logging.GetLogFilePath()
	fileName := logging.GetLogFileName()
	f, err := os.OpenFile(filePath+fileName, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	if err == nil {
		db.SetLogger(log.New(f, "[GORM]", 5))
	}

	// 用于设置最大打开的连接数，默认值为0表示不限制.设置最大的连接数，可以避免并发太高导致连接mysql出现too many connections的错误。
	//db.DB().SetMaxOpenConns(20000)

	// 用于设置闲置的连接数.设置闲置的连接数则当开启的一个连接使用完成后可以放在池里等候下一次使用。
	db.DB().SetMaxIdleConns(0)

	return db
}

func (db *Database) Init() {
	DB = &Database{
		Local: openDB(viper.GetString("local_db.username"),
			viper.GetString("local_db.password"),
			viper.GetString("local_db.addr"),
			viper.GetString("local_db.name"),
		),
	}
}

func (db *Database) Close() {
	DB.Local.Close()
}
