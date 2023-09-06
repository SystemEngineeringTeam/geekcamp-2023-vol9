package model

import (
	"github.com/jinzhu/gorm"
)

// Building is a struct that represent the building table in the database
type Floor struct {
	gorm.Model
	Floor int
	BuildingId int
	Rooms []Room `gorm:"foreignkey:FloorId"`
}
