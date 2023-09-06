package model

import (
	"github.com/jinzhu/gorm"
)

// Building is a struct that represent the building table in the database
type Building struct {
	gorm.Model
	Name string
	Floors []Floor `gorm:"foreignkey:BuildingId"`
}