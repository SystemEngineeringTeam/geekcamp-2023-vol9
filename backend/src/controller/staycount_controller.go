package controller

import (
	"net/http"
	"strconv"
	"time"

	"github.com/SystemEngineeringTeam/geekcamp-2023-vol9/model"
	"github.com/gin-gonic/gin"
)

// @Summary 滞在者数を取得する
// @Description arpscanによって取得したデータを元に滞在者数を取得する、滞在者数は最新のものを取得する
// @Tag StayCount
// @Produce  json
// @Success 200 {object} model.GetStayCountResponseModel
// @Router /api/v1/staycount/get/{building_name} [get]
func StayCountGet(c *gin.Context) {

	// 建物の取得
	building := model.GetStayCount()

	print(building)

	c.JSON(200, gin.H{
		"staycounts": building,
	})
}

// @Summary 滞在者数を登録する
// @Description arpscanによって取得したデータを元に滞在者数を登録する
// @Tag StayCount
// @Produce  json
// @Success 200 {object} model.PostStayCountRequestModel
// @Router /api/v1/staycount/post/{building_name} [POST]
func StayCountPost(c *gin.Context) {

	roomId := c.Param("room_id")

	// postで送信されたデータを取得
	var req model.StayCount
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"type":    "failed",
			"message": "入力されたデータが不正です。",
			"error":   err.Error(),
		})
		return
	}

	// 日本時間に変換
	req.DateTime = req.DateTime.In(time.FixedZone("Asia/Tokyo", 9*60*60)).Add(-9 * time.Hour)

	println(roomId)

	// room_idをint型に変換
	roomIdInt, err := strconv.Atoi(roomId)
	if err == nil {
		req.RoomId = roomIdInt
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"type":    "failed",
			"message": "入力されたroom_idが不正です。",
			"error":   err.Error(),
		})
		return
	}

	// データベースにデータを挿入
	model.StayCountInsert(req)

	// レスポンスを返す
	c.JSON(http.StatusOK, gin.H{
		"type":       "succeeded",
		"room_id":    roomId,
		"stay_count": req.StayCount,
		"date_time":  req.DateTime,
	})
}
