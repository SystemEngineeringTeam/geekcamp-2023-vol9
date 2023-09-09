package router

import (
	"io"
	"os"
    "time"

	"github.com/gin-gonic/gin"
    "github.com/gin-contrib/cors"
    "github.com/SystemEngineeringTeam/geekcamp-2023-vol9/controller"

    swaggerfiles "github.com/swaggo/files"
    ginSwagger "github.com/swaggo/gin-swagger"

)

func Init(){
    // サーバーログの出力先を設定
    gin.DisableConsoleColor()
	f, _ := os.Create("../server.log")
	gin.DefaultWriter = io.MultiWriter(f)

    // ルーティングの設定
	router := gin.Default()

    // ここからCorsの設定
    router.Use(cors.New(cors.Config{
        // アクセスを許可したいアクセス元
        AllowOrigins: []string{
            "https://campus-crowd-monitor.vercel.app/",
            "http://localhost:3000",
        },
        // アクセスを許可したいHTTPメソッド(以下の例だとPUTやDELETEはアクセスできません)
        AllowMethods: []string{
            "POST",
            "GET",
            "OPTIONS",
        },
        // 許可したいHTTPリクエストヘッダ
        AllowHeaders: []string{
            "Access-Control-Allow-Credentials",
            "Access-Control-Allow-Headers",
            "Content-Type",
            "Content-Length",
            "Accept-Encoding",
            "Authorization",
        },
        // cookieなどの情報を必要とするかどうか
        AllowCredentials: true ,
        // preflightリクエストの結果をキャッシュする時間
        MaxAge: 24 * time.Hour,
    }))


    // V1の設定
	v1 := router.Group("/api/v1/")
    v1.GET("/staycount/get/", controller.StayCountGet )
    v1.GET("/congestion/get/", controller.GetCongestionDegree )
    v1.POST("/staycount/post/:room_id", controller.StayCountPost )

    // 下記を追記することで`http://localhost:8080/api/v1/swagger/index.html`を叩くことでswagger uiを開くことができる
    v1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

    // ポート番号の設定
    router.Run(":8080")
}
