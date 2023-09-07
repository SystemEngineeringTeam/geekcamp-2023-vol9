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
}

type GetRoom struct {
	Name string     				`json:"name"`
	RoomId int						`json:"room_id"`
	StayCount int         			`json:"headcount"`
}
