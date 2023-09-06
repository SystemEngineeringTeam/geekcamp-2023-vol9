package controller

import (
	"net/http"

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
			"message": "Invalid request",
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"type": "succeeded",
		"room_id": roomId,
		"stay_count": req.StayCount,
		"date_time": req.DateTime,
	})
}