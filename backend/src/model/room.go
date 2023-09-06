package model

import (
	"github.com/jinzhu/gorm"
)

// Building is a struct that represent the building table in the database	
type Room struct {
	gorm.Model
	Name string
	FloorId int
	StayCounts []StayCount `gorm:"foreignkey:RoomId"`
}
