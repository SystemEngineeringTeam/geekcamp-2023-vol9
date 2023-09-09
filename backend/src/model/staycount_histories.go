package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type History struct {
	gorm.Model
	StayCount [24]int	`json:"staycount"`
	RoomId    int		`json:"-"`
	DateTime  time.Time `json:"time"`
}
