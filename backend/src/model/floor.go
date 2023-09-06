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
