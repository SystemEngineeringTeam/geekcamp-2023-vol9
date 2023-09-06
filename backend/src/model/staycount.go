package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Building is a struct that represent the building table in the database
type StayCount struct {
	gorm.Model
	RoomId int
	StayCount int
	DateTime time.Time
}