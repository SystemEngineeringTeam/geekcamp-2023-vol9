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
// @Success 200 {object} model.GetStayCountHistoryRequestModel
// @Router /api/v1/staycount/history/ [get]
func StayCountHistoriesGet(c *gin.Context) {
	// TODO: しっかりとMySQLからデータを取得する

	histories := model.GetStayCountHistory()

	// RoomIDをキーとしてマップに変換
	// historyMap := make(map[int][24]int)
	// for _, history := range histories {
	// 	historyMap[history.RoomId] = history.StayCount
	// }

	// レスポンスとしてJSONを返す
	c.JSON(http.StatusOK, gin.H{
		"histories": histories,
	})
}
