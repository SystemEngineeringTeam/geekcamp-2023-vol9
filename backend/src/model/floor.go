package model

import (
	"github.com/jinzhu/gorm"
)

// Building is a struct that represent the building table in the database
type Floor struct {
	gorm.Model
	Floor int				`json:"floor"`
	BuildingId int			`json:"building_id"`
	Rooms []Room 			`gorm:"foreignkey:FloorId"`
}

type GetStayCountFloorModel struct {
	Floor int							`json:"floor"`
	Rooms []GetStayCountRoomModel 		`json:"rooms" `
}

type GetCongestionFloorModel struct {
	Floor int							`json:"floor"`
	Rooms []GetCongestionRoomModel 		`json:"rooms" `
}
