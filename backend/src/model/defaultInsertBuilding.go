package model

import (
	"fmt"
	"log"
	"time"
)

// TODO : 一番最初のデータベースに入れる部分は、最終的にSQLにしてDocker起動時に反映するように変更する。

func BuildingDefaultInsert(){
	building := []Building{
		{
			Name: "4号館別館",
			Floors: []Floor{
				{
					Floor: 1,
					Rooms: []Room{
						{
							Name: "1F",
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
							Name: "2F",
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
							Name: "3F",
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
							Name: "4F",
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
		},	
		{
			Name: "4号館",	
			Floors: []Floor{
				{
					Floor: 1,
					Rooms: []Room{
						{
							Name: "1F",
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
							Name: "2F",
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
							Name: "3F",
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
							Name: "4F",
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
		},
		{
			Name: "1号館",
			Floors: []Floor{
				{
					Floor: 1,
					Rooms: []Room{
						{
							Name: "1F",
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
							Name: "2F",
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
							Name: "3F",
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
							Name: "4F",
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
					Floor: 5,
					Rooms: []Room{
						{
							Name: "501",
							StayCounts: []StayCount{
								{
									StayCount: 0,
									DateTime: time.Now(),
								},
							},
						},
						{
							Name: "502",
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
					Floor: 6,
					Rooms: []Room{
						{
							Name: "6F",
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
					Floor: 7,
					Rooms: []Room{
						{
							Name: "7F",
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
		},
		{
			Name: "10号館",
			Floors: []Floor{
				{
					Floor: 1,
					Rooms: []Room{
						{
							Name: "1F",
							StayCounts: []StayCount{
								{
									StayCount: 0,
									DateTime: time.Now(),
								},
							},
						},
						{
							Name: "G2110",
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
							Name: "2F",
							StayCounts: []StayCount{
								{
									StayCount: 0,
									DateTime: time.Now(),
								},
							},
						},
						{
							Name: "G2210",
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
							Name: "3F",
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
							Name: "4F",
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
					Floor: 5,
					Rooms: []Room{
						{
							Name: "5F",
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
					Floor: 6,
					Rooms: []Room{
						{
							Name: "6F",
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
					Floor: 7,
					Rooms: []Room{
						{
							Name: "7F",
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
		},
		{
			Name: "セントラル",
			Floors: []Floor{
				{
					Floor: 1,
					Rooms: []Room{
						{
							Name: "セントラル",
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
		},
		{
			Name: "AITプラザ",
			Floors: []Floor{
				{
					Floor: 1,
					Rooms: []Room{
						{
							Name: "AITプラザ",
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
		},
		{
			Name: "愛和会館",
			Floors: []Floor{
				{
					Floor: 1,
					Rooms: []Room{
						{
							Name: "1F",
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
							Name: "2F",
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
		},
	}

	result := db.Create(&building)
	if result.Error != nil {
		log.Fatal(result.Error)
	}

	println("登録したよ")
	for _ , building := range building {
		fmt.Println("Inserted a record successfully:" , building.Name)
	}
}