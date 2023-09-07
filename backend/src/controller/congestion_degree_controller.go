package controller

import (

	"github.com/SystemEngineeringTeam/geekcamp-2023-vol9/model"
	"github.com/gin-gonic/gin"
)

func GetCongestionDegree (c *gin.Context){

	buildingName := c.Param("building_name")

	// 建物の取得
	building := model.GetCongestionDegree(buildingName)

	c.JSON(200, gin.H{
		"congestions": building,
	})
}