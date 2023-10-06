package model

import (
	"github.com/jinzhu/gorm"
)

// Building is a struct that represent the building table in the database
type Room struct {
	gorm.Model
	Name string     		`json:"name"`
	FloorId int				`json:"floor_id"`
	StayCounts []StayCount 	`gorm:"foreignkey:RoomId"`
	// History []History 		`gorm:"foreignkey:RoomId"`
}

type GetStayCountRoomModel struct {
	Name string     				`json:"name"`
	RoomId int						`json:"id"`
	StayCount int         			`json:"staycount"`
	MaxStaycount int				`json:"maxStaycount"`
}

type GetCongestionRoomModel struct {
	Name string     					`json:"name"`
	RoomId int							`json:"id"`
	Congestion float64       			`json:"congestion"`
}
