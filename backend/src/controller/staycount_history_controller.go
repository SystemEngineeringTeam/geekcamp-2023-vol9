package controller

import (
	// "net/http"
	// "strconv"

	"net/http"
	"strconv"
	"time"

	"github.com/SystemEngineeringTeam/geekcamp-2023-vol9/model"
	"github.com/gin-gonic/gin"
)

// @Summary 滞在者数の履歴を取得する（今日だけ）
// @Description 滞在者数の履歴を取得する。これは24時間分のデータを取得する。
// @Tag StayCount
// @Produce  json
// @Success 200 {object} model.GetStayCountHistoryRequestModel
// @Router /api/v1/staycount/history/ [get]
func StayCountHistoriesGet(c *gin.Context) {
	// TODO: しっかりとMySQLからデータを取得する

	histories := model.GetStayCountHistory()

	// レスポンスとしてJSONを返す
	c.JSON(http.StatusOK, gin.H{
		"histories": histories,
	})
}

// @Summary 滞在者数の履歴を取得する。部屋指定かつ日付指定
// @Description 滞在者数の履歴を取得する。これは24時間分のデータを取得する。
// @Tag StayCount
// @Produce  json
// @Success 200 {object} model.GetStayCountHistoryModel
// @Param date query string true "日付データを記載する 例: 2021-10-01"
// @Router /api/v1/staycount/history/{room_id} [get]
func StayCountHistoriesGetByRoomIdAndDate(c *gin.Context) {

	roomId := c.Param("room_id")
	date := c.Query("date")

	// Time.timeに変換
	date_time, err := time.Parse("2006-01-02", date)
	if err != nil {
		panic(err)
	}

	//roomIdをint型に変換
	roomIdInt, err := strconv.Atoi(roomId)
	if err != nil {
		panic(err)
	}

	history := model.GetStayCountHistoryByRoomIdAndDate(date_time, roomIdInt)

	// レスポンスとしてJSONを返す
	c.JSON(http.StatusOK, gin.H{
		"histories": history,
	})
}
