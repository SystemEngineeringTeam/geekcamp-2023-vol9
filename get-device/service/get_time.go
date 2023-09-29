package service

import (
	"fmt"
	"time"
)

func GetTime() string {
	jst, err := time.LoadLocation(("Asia/Tokyo"))
	if err != nil {
		panic(err)
	}
	t := time.Now().In(jst)
	time := t.Format("2006-01-02T15:04:00.000Z")
	fmt.Println(time)
	return (time)
}
