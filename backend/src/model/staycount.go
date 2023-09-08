package model

import (
	"fmt"
	"log"
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

func StayCountInsert(staycount StayCount) {
	result := db.Create(&staycount)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	fmt.Println("Inserted a record successfully:" , staycount.RoomId ,  staycount.StayCount, staycount.DateTime)
}