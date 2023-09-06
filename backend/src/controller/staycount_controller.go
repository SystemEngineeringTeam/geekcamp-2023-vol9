package controller

import (
	"net/http"
	"strconv"

	"github.com/SystemEngineeringTeam/geekcamp-2023-vol9/model"
	"github.com/gin-gonic/gin"
)

func StayCountGet(c *gin.Context){

	buildingName := c.Param("building_name")

	c.JSON(200, gin.H{
		"type": "succeeded",
		"building_name": buildingName,
	})
}

func StayCountPost(c *gin.Context){
	
	roomId := c.Param("room_id")

	// postで送信されたデータを取得
	var req model.StayCount
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"type": "failed",
			"message": "入力されたデータが不正です。",
			"error": err.Error(),
		})
		return
	}

	println(roomId)

	// room_idをint型に変換
	roomIdInt, err := strconv.Atoi(roomId)
	if err == nil {
		req.RoomId = roomIdInt
	}else{
		c.JSON(http.StatusBadRequest, gin.H{
			"type": "failed",
			"message": "入力されたroom_idが不正です。",
			"error": err.Error(),
		})
		return
	}

	// データベースにデータを挿入
	model.StayCountInsert(req)

	// レスポンスを返す
	c.JSON(http.StatusOK, gin.H{
		"type": "succeeded",
		"room_id": roomId,
		"stay_count": req.StayCount,
		"date_time": req.DateTime,
	})
}