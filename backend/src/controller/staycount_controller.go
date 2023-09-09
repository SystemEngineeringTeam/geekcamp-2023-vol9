package controller

import (
	"net/http"
	"strconv"

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
	// building := model.GetStayCount()
	// println(building.Floors[0].Rooms[0].StayCounts[0].StayCount)

	// 嘘データを作成

	building := []model.GetStayCountBuildingModel{
		{
			Name: "1号館",
			Floors: []model.GetStayCountFloorModel{
				{
					Floor: 2,
					Rooms: []model.GetStayCountRoomModel{
						{
							Name:      "アメリカフェ",
							RoomId:    0,
							StayCount: 3,
						},
					},
				},
				{
					Floor: 1,
					Rooms: []model.GetStayCountRoomModel{
						{
							Name:      "キャリアセンター",
							RoomId:    1,
							StayCount: 0,
						},
					},
				},
				{
					Floor: 3,
					Rooms: []model.GetStayCountRoomModel{
						{
							Name:      "301",
							RoomId:    2,
							StayCount: 5,
						},
						{
							Name:      "302",
							RoomId:    3,
							StayCount: 2,
						},
					},
				},
				{
					Floor: 4,
					Rooms: []model.GetStayCountRoomModel{
						{
							Name:      "401",
							RoomId:    4,
							StayCount: 0,
						},
						{
							Name:      "402",
							RoomId:    5,
							StayCount: 0,
						},
					},
				},
				{
					Floor: 5,
					Rooms: []model.GetStayCountRoomModel{
						{
							Name:      "501",
							RoomId:    6,
							StayCount: 0,
						},
						{
							Name:      "502",
							RoomId:    7,
							StayCount: 0,
						},
					},
				},
				{
					Floor: 6,
					Rooms: []model.GetStayCountRoomModel{
						{
							Name:      "601",
							RoomId:    8,
							StayCount: 0,
						},
						{
							Name:      "602",
							RoomId:    9,
							StayCount: 0,
						},
					},
				},
				{
					Floor: 7,
					Rooms: []model.GetStayCountRoomModel{
						{
							Name:      "701",
							RoomId:    10,
							StayCount: 0,
						},
						{
							Name:      "702",
							RoomId:    11,
							StayCount: 0,
						},
						{
							Name:      "鳥居教授",
							RoomId:    12,
							StayCount: 0,
						},
						{
							Name:      "鳥居研",
							RoomId:    13,
							StayCount: 0,
						},
					},
				},
			},
		},
		{
			Name: "10号館",
			Floors: []model.GetStayCountFloorModel{
				{
					Floor: 1,
					Rooms: []model.GetStayCountRoomModel{},
				},
				{
					Floor: 2,
					Rooms: []model.GetStayCountRoomModel{
						{
							Name:      "G2201",
							RoomId:    14,
							StayCount: 0,
						},
						{
							Name:      "G2202",
							RoomId:    15,
							StayCount: 0,
						},
						{
							Name:      "G2203",
							RoomId:    16,
							StayCount: 0,
						},
						{
							Name:      "G2204",
							RoomId:    17,
							StayCount: 0,
						},
					},
				},
				{
					Floor: 3,
					Rooms: []model.GetStayCountRoomModel{
						{
							Name:      "G2305",
							RoomId:    18,
							StayCount: 0,
						},
						{
							Name:      "G2306",
							RoomId:    19,
							StayCount: 0,
						},
						{
							Name:      "G2307",
							RoomId:    20,
							StayCount: 0,
						},
						{
							Name:      "G2308",
							RoomId:    21,
							StayCount: 0,
						},
					},
				},
				{
					Floor: 4,
					Rooms: []model.GetStayCountRoomModel{
						{
							Name:      "G2409",
							RoomId:    22,
							StayCount: 0,
						},
						{
							Name:      "G2410",
							RoomId:    23,
							StayCount: 0,
						},
						{
							Name:      "G2411",
							RoomId:    24,
							StayCount: 0,
						},
						{
							Name:      "G2412",
							RoomId:    25,
							StayCount: 0,
						},
					},
				},
				{
					Floor: 5,
					Rooms: []model.GetStayCountRoomModel{
						{
							Name:      "G2513",
							RoomId:    26,
							StayCount: 0,
						},
						{
							Name:      "G2514",
							RoomId:    27,
							StayCount: 0,
						},
						{
							Name:      "G2515",
							RoomId:    28,
							StayCount: 0,
						},
						{
							Name:      "G2516",
							RoomId:    29,
							StayCount: 0,
						},
					},
				},
				{
					Floor: 6,
					Rooms: []model.GetStayCountRoomModel{
						{
							Name:      "G2617",
							RoomId:    30,
							StayCount: 0,
						},
						{
							Name:      "G2618",
							RoomId:    31,
							StayCount: 0,
						},
						{
							Name:      "G2619",
							RoomId:    32,
							StayCount: 0,
						},
						{
							Name:      "G2620",
							RoomId:    33,
							StayCount: 0,
						},
					},
				},
				{
					Floor: 7,
					Rooms: []model.GetStayCountRoomModel{
						{
							Name:      "G2621",
							RoomId:    34,
							StayCount: 0,
						},
						{
							Name:      "G2622",
							RoomId:    35,
							StayCount: 0,
						},
						{
							Name:      "G2623",
							RoomId:    36,
							StayCount: 0,
						},
						{
							Name:      "G2624",
							RoomId:    37,
							StayCount: 0,
						},
					},
				},
			},
		},
		{
			Name: "4号館",
			Floors: []model.GetStayCountFloorModel{
				{
					Floor: 1,
					Rooms: []model.GetStayCountRoomModel{
						{
							Name:      "内種教授",
							RoomId:    38,
							StayCount: 0,
						},
						{
							Name:      "内種研",
							RoomId:    39,
							StayCount: 1,
						},
						{
							Name:      "玉森准教授",
							RoomId:    40,
							StayCount: 0,
						},
						{
							Name:      "玉森研",
							RoomId:    41,
							StayCount: 3,
						},
					},
				},
				{
					Floor: 2,
					Rooms: []model.GetStayCountRoomModel{
						{
							Name:      "小野木教授",
							RoomId:    42,
							StayCount: 0,
						},
						{
							Name:      "小野木研",
							RoomId:    43,
							StayCount: 0,
						},
						{
							Name:      "小栗講師",
							RoomId:    44,
							StayCount: 0,
						},
						{
							Name:      "塚田教授",
							RoomId:    45,
							StayCount: 0,
						},
						{
							Name:      "塚田研",
							RoomId:    46,
							StayCount: 0,
						},
						{
							Name:      "森本教授",
							RoomId:    47,
							StayCount: 0,
						},
						{
							Name:      "森本研",
							RoomId:    48,
							StayCount: 0,
						},
					},
				},
				{
					Floor: 3,
					Rooms: []model.GetStayCountRoomModel{
						{
							Name:      "伊藤雅教授",
							RoomId:    49,
							StayCount: 0,
						},
						{
							Name:      "伊藤雅研",
							RoomId:    50,
							StayCount: 0,
						},
						{
							Name:      "中村教授",
							RoomId:    51,
							StayCount: 0,
						},
						{
							Name:      "中村研",
							RoomId:    52,
							StayCount: 1,
						},
						{
							Name:      "菱田教授",
							RoomId:    53,
							StayCount: 0,
						},
						{
							Name:      "菱田研",
							RoomId:    54,
							StayCount: 0,
						},
					},
				},
				{
					Floor: 4,
					Rooms: []model.GetStayCountRoomModel{
						{
							Name:      "401",
							RoomId:    55,
							StayCount: 0,
						},
						{
							Name:      "402",
							RoomId:    56,
							StayCount: 0,
						},
						{
							Name:      "403",
							RoomId:    57,
							StayCount: 0,
						},
						{
							Name:      "404",
							RoomId:    58,
							StayCount: 0,
						},
					},
				},
			},
		},
		{
			Name: "4号館別館",
			Floors: []model.GetStayCountFloorModel{
				{
					Floor: 1,
					Rooms: []model.GetStayCountRoomModel{
						{
							Name:      "事務室",
							RoomId:    59,
							StayCount: 0,
						},
						{
							Name:      "梶准教授",
							RoomId:    60,
							StayCount: 0,
						},
						{
							Name:      "梶研",
							RoomId:    61,
							StayCount: 5,
						},
					},
				},
				{
					Floor: 2,
					Rooms: []model.GetStayCountRoomModel{
						{
							Name:      "伊藤暢浩教授",
							RoomId:    62,
							StayCount: 0,
						},
						{
							Name:      "伊藤暢浩研",
							RoomId:    63,
							StayCount: 3,
						},
						{
							Name:      "河辺教授",
							RoomId:    64,
							StayCount: 0,
						},
						{
							Name:      "河辺研",
							RoomId:    65,
							StayCount: 0,
						},
					},
				},
				{
					Floor: 3,
					Rooms: []model.GetStayCountRoomModel{
						{
							Name:      "内藤教授",
							RoomId:    66,
							StayCount: 0,
						},
						{
							Name:      "内藤研",
							RoomId:    67,
							StayCount: 0,
						},
						{
							Name:      "松河教授",
							RoomId:    68,
							StayCount: 0,
						},
						{
							Name:      "松河研",
							RoomId:    69,
							StayCount: 0,
						},
						{
							Name:      "シス研",
							RoomId:    70,
							StayCount: 6,
						},
					},
				},
				{
					Floor: 4,
					Rooms: []model.GetStayCountRoomModel{
						{
							Name:      "澤野准教授",
							RoomId:    71,
							StayCount: 0,
						},
						{
							Name:      "澤野研",
							RoomId:    72,
							StayCount: 0,
						},
						{
							Name:      "水野教授",
							RoomId:    73,
							StayCount: 0,
						},
						{
							Name:      "水野研",
							RoomId:    74,
							StayCount: 0,
						},
						{
							Name:      "山本教授",
							RoomId:    75,
							StayCount: 0,
						},
						{
							Name:      "山本研",
							RoomId:    76,
							StayCount: 0,
						},
					},
				},
			},
		},
	}

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
