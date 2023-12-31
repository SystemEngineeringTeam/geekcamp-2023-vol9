package main

import (
	"github.com/SystemEngineeringTeam/geekcamp-2023-vol9/model"
	"github.com/SystemEngineeringTeam/geekcamp-2023-vol9/router"
    _ "github.com/SystemEngineeringTeam/geekcamp-2023-vol9/docs"
)


// @title CampusCrowdMonitor
// @version 1.0
// @description このswaggerはCampusCrowdMonitorのapiです
func main() {
    // テーブル作成とDB接続
    model.CreateAllTable()
    // ルーティングの設定＋サーバー起動
    router.Init()
}