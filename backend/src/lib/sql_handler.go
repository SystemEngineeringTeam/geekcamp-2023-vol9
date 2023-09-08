package lib

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// MYSQLの接続情報
func SqlConnect() (database *gorm.DB) {
	var db *gorm.DB
	var err error
	// TODO: 環境変数から取得するようにする
	dsn := "root:admin@tcp(127.0.0.1:3309)/koukaten2023?charset=utf8mb4&parseTime=true&loc=Asia%2FTokyo"
	dialector := mysql.Open(dsn)

	if db, err = gorm.Open(dialector); err != nil {
		db = connect(dialector, 10)
	}
	fmt.Println("db connected!!")

	return db
}

// SQLの接続
func connect(dialector gorm.Dialector, count uint) *gorm.DB {
	var err error
	var db *gorm.DB
	if db, err = gorm.Open(dialector); err != nil {
		if count > 1 {
			time.Sleep(time.Second * 2)
			count--
			fmt.Printf("retry... count:%v\n", count)
			connect(dialector, count)
		}
		panic(err.Error())
	}
	return (db)
}