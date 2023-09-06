package model

import (
    "github.com/SystemEngineeringTeam/geekcamp-2023-vol9/lib"
)

// MYSQLの接続情報
var db = lib.SqlConnect()

// テーブル作成
func CreateAllTable(){
    db.AutoMigrate(&Building{})
    db.AutoMigrate(&Floor{})
    db.AutoMigrate(&Room{})
    db.AutoMigrate(&StayCount{})
}
