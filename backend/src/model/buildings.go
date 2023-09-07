package model

import (
	"github.com/jinzhu/gorm"
)

// Building is a struct that represent the building table in the database
type Building struct {
	gorm.Model
	Name string			`json:"name"`
	Floors []Floor 		`gorm:"foreignkey:BuildingId"`
}

type GetBuilding struct {
	Name string				`json:"name"`
	Floors []GetFloor 		`json:"floors"`
}

func GetStayCount(buildingName string) GetBuilding {

	// TODO: いっぱいSQLを飛ばしてしまうため、改善の余地あり
	// DBの結合がヒントらしい
	// https://syoblog.com/golang-gorm-join/
	// https://wa3.i-3-i.info/word15318.html
	// 部室にある応用情報の本のDBのセクションに載ってるのがわかりやすいらしい。

	// Jsonに変換するための構造体
	building := Building{}
	req := GetBuilding{}

	db.Preload("Floors.Rooms.StayCounts").
		Where("name = ?", buildingName).
		First(&building)

	// reqに詰め替える
	reqFloars := []GetFloor{}
	for _, floor := range building.Floors {

		reqRooms := []GetRoom{}
		for _, room := range floor.Rooms {
			reqRooms = append(reqRooms, GetRoom{
				Name: room.Name,
				RoomId: int(room.ID),
				StayCount: room.StayCounts[len(room.StayCounts)-1].StayCount,
			})
		}

		reqFloars = append(reqFloars, GetFloor{
			Floor: floor.Floor,
			Rooms: reqRooms,
		})
	}
	req.Floors = reqFloars
	req.Name = building.Name

	return req
}
