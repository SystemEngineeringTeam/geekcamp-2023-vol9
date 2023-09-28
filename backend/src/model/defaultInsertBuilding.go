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
		},	
		{
			Name: "4号館",	
			Floors: []Floor{
				{
					Floor: 1,
					Rooms: []Room{
						{
							Name: "4号館1F",
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
							Name: "4号館2F",
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
							Name: "4号館3F",
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
							Name: "4号館4F",
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
							Name: "1号館1F",
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
							Name: "1号館2F",
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
							Name: "1号館3F",
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
							Name: "1号館4F",
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
							Name: "1号館5F-501",
							StayCounts: []StayCount{
								{
									StayCount: 0,
									DateTime: time.Now(),
								},
							},
						},
						{
							Name: "1号館5F-502",
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
							Name: "1号館6F",
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
							Name: "1号館7F",
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
							Name: "10号館1F",
							StayCounts: []StayCount{
								{
									StayCount: 0,
									DateTime: time.Now(),
								},
							},
						},
						{
							Name: "10号館1F-G2110",
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
							Name: "10号館2F",
							StayCounts: []StayCount{
								{
									StayCount: 0,
									DateTime: time.Now(),
								},
							},
						},
						{
							Name: "10号館2F-G2210",
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
							Name: "10号館3F",
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
							Name: "10号館4F",
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
							Name: "10号館5F",
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
							Name: "10号館6F",
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
							Name: "10号館7F",
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
							Name: "愛和会館1F",
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
							Name: "愛和会館2F",
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