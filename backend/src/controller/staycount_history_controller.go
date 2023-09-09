package controller

import (
	// "net/http"
	// "strconv"

	"net/http"

	"github.com/SystemEngineeringTeam/geekcamp-2023-vol9/model"
	"github.com/gin-gonic/gin"
)

// @Summary 滞在者数の履歴を取得する（今日だけ）
// @Description 滞在者数の履歴を取得する。これは24時間分のデータを取得する。
// @Tag StayCount
// @Produce  json
// @Success 200 {object} model.GetStayCountResponseModel
// @Router /api/v1/staycount/history/ [get]
func StayCountHistoriesGet(c *gin.Context) {
	// TODO: しっかりとMySQLからデータを取得する

	histories := []model.History{
		{
			StayCount: [24]int{
				0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 3, 2, 0, 3, 3, 4, 0, 0, 0, 0, 0, 0, 0,
			},
			RoomId: 0,
		},
		{
			StayCount: [24]int{
				0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			},
			RoomId: 1,
		},
		{
			StayCount: [24]int{
				0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 3, 2, 0, 3, 4, 4, 0, 0, 0, 0, 0, 0, 0,
			},
			RoomId: 2,
		},
		{
			StayCount: [24]int{
				0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2, 2, 2, 2, 3, 2, 2, 0, 0, 0, 0, 0, 0, 0,
			},
			RoomId: 3,
		},
		{
			StayCount: [24]int{
				0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			},
			RoomId: 13,
		},
		{
			StayCount: [24]int{
				0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2, 2, 2, 0, 0, 0, 0, 0, 0, 0,
			},
			RoomId: 39,
		},
		{
			StayCount: [24]int{
				0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2, 2, 2, 2, 3, 3, 3, 0, 0, 0, 0, 0, 0, 0,
			},
			RoomId: 41,
		},
		{
			StayCount: [24]int{
				0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			},
			RoomId: 44,
		},
		{
			StayCount: [24]int{
				0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0,
			},
			RoomId: 45,
		},
		{
			StayCount: [24]int{
				2, 4, 4, 2, 2, 2, 2, 3, 4, 4, 4, 3, 4, 8, 8, 8, 5, 0, 0, 0, 0, 0, 0, 0,
			},
			RoomId: 61,
		},
		{
			StayCount: [24]int{
				0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 2, 2, 2, 3, 3, 0, 0, 0, 0, 0, 0, 0,
			},
			RoomId: 63,
		},
		{
			StayCount: [24]int{
				3, 0, 0, 3, 3, 3, 3, 3, 5, 5, 5, 3, 6, 6, 6, 6, 5, 0, 0, 0, 0, 0, 0, 0,
			},
			RoomId: 70,
		},
		{
			StayCount: [24]int{
				0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			},
			RoomId: 71,
		},
	}

	// RoomIDをキーとしてマップに変換
	historyMap := make(map[int][24]int)
	for _, history := range histories {
		historyMap[history.RoomId] = history.StayCount
	}

	// レスポンスとしてJSONを返す
	c.JSON(http.StatusOK, gin.H{
		"histories": historyMap,
	})
}
