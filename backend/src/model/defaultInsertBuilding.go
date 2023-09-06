package model

import (
	"fmt"
	"log"
	"time"
)

func BuildingDefaultInsert(){
	building := Building{
		Name: "四号館別館",
		Floors: []Floor{
			{
				Floor: 1,
				Rooms: []Room{
					{
						Name: "4号館別館1F",
						StayCounts: []StayCount{
							{
								StayCount: 0,
								DateTime: time.Now(),
							},
						},
					},
				},
			},
			{
				Floor: 2,
				Rooms: []Room{
					{
						Name: "4号館別館2F",
						StayCounts: []StayCount{
							{
								StayCount: 0,
								DateTime: time.Now(),
							},
						},
					},
				},
			},
			{
				Floor: 3,
				Rooms: []Room{
					{
						Name: "4号館別館3F",
						StayCounts: []StayCount{
							{
								StayCount: 0,
								DateTime: time.Now(),
							},
						},
					},
				},
			},
			{
				Floor: 4,
				Rooms: []Room{
					{
						Name: "4号館別館4F",
						StayCounts: []StayCount{
							{
								StayCount: 0,
								DateTime: time.Now(),
							},
						},
					},
				},
			},
		},				
	}

	result := db.Create(&building)
	if result.Error != nil {
		log.Fatal(result.Error)
	}

	println("登録したよ")
	fmt.Println("Inserted a record successfully:" , building.Name)
}