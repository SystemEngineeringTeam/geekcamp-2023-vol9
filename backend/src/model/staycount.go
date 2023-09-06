package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Building is a struct that represent the building table in the database
type StayCount struct {
	gorm.Model
	RoomId int				`json:"room_id"`
	StayCount int         	`json:"headcount"`
	DateTime time.Time		`json:"time"`
}