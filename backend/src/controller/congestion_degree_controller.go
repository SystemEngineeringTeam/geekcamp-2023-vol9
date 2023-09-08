package controller

import (

	"github.com/SystemEngineeringTeam/geekcamp-2023-vol9/model"
	"github.com/gin-gonic/gin"
)

// @Summary 混雑度を取得する
// @Description 混雑度は、今までの滞在者数の最大値 / 現在の滞在者数で求める。
// @Tag StayCount
// @Produce  json    
// @Success 200 {object} model.GetCongestionResponseModel
// @Router /api/v1/congestion/get/{building_name} [get]
func GetCongestionDegree (c *gin.Context){

	// 建物の取得
	building := model.GetCongestionDegree()

	c.JSON(200, gin.H{
		"congestions": building,
	})
}