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