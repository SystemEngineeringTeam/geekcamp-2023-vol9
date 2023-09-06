package router

import (
	"io"
	"os"

	"github.com/gin-gonic/gin"
    "github.com/SystemEngineeringTeam/geekcamp-2023-vol9/controller"
)

func Init(){
    // サーバーログの出力先を設定
    gin.DisableConsoleColor()
	f, _ := os.Create("../server.log")
	gin.DefaultWriter = io.MultiWriter(f)

    // ルーティングの設定
	router := gin.Default()

    // V1の設定
	v1 := router.Group("/api/v1/")
    v1.GET("/staycount/get/:building_name", controller.StayCountGet )
    v1.POST("/staycount/post/:room_id", controller.StayCountPost )

    // ポート番号の設定
    router.Run(":8080")
}
