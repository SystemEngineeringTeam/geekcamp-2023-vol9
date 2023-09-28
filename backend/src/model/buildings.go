package model

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

// Building is a struct that represent the building table in the database
type Building struct {
	gorm.Model
	Name   string  `json:"building"`
	Floors []Floor `gorm:"foreignkey:BuildingId"`
}

type GetStayCountBuildingModel struct {
	Name   string                   `json:"building"`
	Floors []GetStayCountFloorModel `json:"floors"`
}

type GetCongestionBuildingModel struct {
	Name   string                    `json:"building"`
	Floors []GetCongestionFloorModel `json:"floors"`
}


func GetStayCount() []GetStayCountBuildingModel {

	// TODO: いっぱいSQLを飛ばしてしまうため、改善の余地あり
	// DBの結合がヒントらしい
	// https://syoblog.com/golang-gorm-join/
	// https://wa3.i-3-i.info/word15318.html
	// 部室にある応用情報の本のDBのセクションに載ってるのがわかりやすいらしい。

	// Jsonに変換するための構造体
	building := []Building{}

	// TODO: StayCountが大量にあるときに、最新のものだけ取得するようにする
	db.Preload("Floors.Rooms.StayCounts").
		Find(&building)

	fmt.Println(building)

	// reqに詰め替える
	req := []GetStayCountBuildingModel{}
	for _, building := range building {
		reqFloars := []GetStayCountFloorModel{}
		for _, floor := range building.Floors {

			reqRooms := []GetStayCountRoomModel{}
			for _, room := range floor.Rooms {
				reqRooms = append(reqRooms, GetStayCountRoomModel{
					Name:      room.Name,
					RoomId:    int(room.ID),
					StayCount: room.StayCounts[len(room.StayCounts)-1].StayCount,
				})
			}

			reqFloars = append(reqFloars, GetStayCountFloorModel{
				Floor: floor.Floor,
				Rooms: reqRooms,
			})
		}

		req = append(req, GetStayCountBuildingModel{
			Name:   building.Name,
			Floors: reqFloars,
		})
	}

	return req

}

func GetCongestionDegree() []GetCongestionBuildingModel {
	// 建物の取得
	building := []Building{}
	req := []GetCongestionBuildingModel{}
	db.Find(&building)

	for _, building := range building {

		// 建物に紐づく階の取得
		floors := []Floor{}
		reqFloars := []GetCongestionFloorModel{}
		db.Where("building_id = ?", building.ID).Find(&floors)

		for _, floor := range floors {

			// 階に紐づく部屋の取得
			rooms := []Room{}
			reqRooms := []GetCongestionRoomModel{}
			db.Where("floor_id = ?", floor.ID).Find(&rooms)

			for roomIndex, room := range rooms {

				// 部屋に紐づく滞在人数の取得
				maxStayCount := StayCount{}
				stayCount := []StayCount{}
				db.Where("room_id = ?", room.ID).Order("stay_count DESC").First(&maxStayCount)
				db.Where("room_id = ?", room.ID).Last(&stayCount)

				// 混雑度の計算
				congestion := 0.0
				if maxStayCount.StayCount == 0 {
					congestion = 0
				} else {
					congestion = float64(stayCount[0].StayCount) / float64(maxStayCount.StayCount) * 100
					rooms[roomIndex].StayCounts = stayCount
				}

				// 確認
				println(building.Name, floor.Floor, room.Name, congestion)

				// reqに詰め替える
				reqRooms = append(reqRooms, GetCongestionRoomModel{
					Name:       room.Name,
					RoomId:     int(room.ID),
					Congestion: congestion,
				})
			}
			reqFloars = append(reqFloars, GetCongestionFloorModel{
				Floor: floor.Floor,
				Rooms: reqRooms,
			})
		}

		req = append(req, GetCongestionBuildingModel{
			Name:   building.Name,
			Floors: reqFloars,
		})
	}

	return req

}
