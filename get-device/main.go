package main

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"

	"github.com/SystemEngineeringTeam/geekcamp-2023-vol9/get-device/service"
)

type netowork struct {
	RoomId int
	Ip     string
	Mask   int
}

func main() {
	net := getIpAndMask()
	// ip := "192.168.0.3"
	// mask := 24
	time := service.GetTime()
	for _, v := range net {
		// fmt.Println(v)
		c := service.Ping(v.Ip, v.Mask)
		// fmt.Println(c)
		service.PostCount(v.RoomId, c, time)
	}
}

func getIpAndMask() []netowork {
	file, err := os.Open("./network.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	r := csv.NewReader(file)
	rows, err := r.ReadAll() // csvを一度に全て読み込む
	if err != nil {
		log.Fatal(err)
	}
	var net []netowork
	for i, v := range rows {
		if i == 0 {
			continue
		}
		id, _ := strconv.Atoi(v[0])
		mask, _ := strconv.Atoi(v[2])
		n := netowork{
			RoomId: id,
			Ip:     v[1],
			Mask:   mask,
		}
		net = append(net, n)
	}
	return net
}
