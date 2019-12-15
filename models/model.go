package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
	"lemon/utils/logging"
	"log"
	"os"
	"time"
)

type BaseModel struct {
	Id        uint64     `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"-"`
	CreatedAt *time.Time `gorm:"column:create_at" json:"-"`
	UpdatedAt *time.Time `gorm:"column:update_at" json:"-"`
}

type Database struct {
	Default *gorm.DB
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
		logging.Errorf("Database connection failed. Database name: %s, Error info: %s", name, err)
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
		Default: openDB(viper.GetString("default_db.username"),
			viper.GetString("default_db.password"),
			viper.GetString("default_db.addr"),
			viper.GetString("default_db.name"),
		),
	}
}

func (db *Database) Close() {
	DB.Default.Close()
}
