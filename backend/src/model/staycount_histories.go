package model

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

type History struct {
	gorm.Model
	StayCount [24]int   `json:"staycount"`
	RoomId    int       `json:"-"`
	DateTime  time.Time `json:"time"`
}



func GetStayCountHistory() map[int][]float64 {

	// 今日の日付を取得
	today := time.Now().Format("2006-01-02")

	// 建物の取得
	// building := []Building{}
	// 日付を今日のものだけにする
	// StayCountモデルから今日のデータを取得
	stayCounts := []StayCount{}
	if err := db.Where("date_time >= ? AND date_time < ?", today, today+" 23:59:59").Find(&stayCounts).Error; err != nil {
		fmt.Println("データの取得に失敗しました:", err)
		return nil
	}

	// 時間ごとのデータを格納するためのマップ
	roomHourlyData := make(map[int][]float64)

	// 部屋ごとに時間帯ごとにデータを集計
	for _, sc := range stayCounts {
		roomId := sc.RoomId
		hourlyTime := sc.DateTime.Truncate(time.Hour) // 1時間ごとに区切る
		if _, ok := roomHourlyData[roomId]; !ok {
			roomHourlyData[roomId] = make([]float64, 24) // 24時間分のスライスを用意
		}
		// 時間帯に対応するIndexを計算
		hourIndex := hourlyTime.Hour()
		if roomHourlyData[roomId][hourIndex] == 0 {
			roomHourlyData[roomId][hourIndex] += float64(sc.StayCount)
		} else {
			roomHourlyData[roomId][hourIndex] += float64(sc.StayCount)
			roomHourlyData[roomId][hourIndex] /= 2
		}

	}

	// 最終的なJSONを作成
	result := make(map[int][]float64)
	for roomId, hourlyData := range roomHourlyData {
		result[roomId] = hourlyData
	}

	for _, sc := range result {
		fmt.Println(sc)
	}

	return result
}

func GetStayCountHistoryByRoomIdAndDate( date time.Time , roomId int )  map[int][]float64{
	stayCounts := []StayCount{}
	date = date.In(time.FixedZone("Asia/Tokyo", 9*60*60))
	db.Where("room_id = ? AND date_time >= ? AND date_time < ?", roomId, date, date.Add(24*time.Hour)).Find(&stayCounts)


	// 時間ごとのデータを格納するためのマップ
	roomHourlyData := make(map[int][]float64)

	// 部屋ごとに時間帯ごとにデータを集計
	for _, sc := range stayCounts {
		roomId := sc.RoomId
		hourlyTime := sc.DateTime.Truncate(time.Hour) // 1時間ごとに区切る
		if _, ok := roomHourlyData[roomId]; !ok {
			roomHourlyData[roomId] = make([]float64, 24) // 24時間分のスライスを用意
		}
		// 時間帯に対応するIndexを計算
		hourIndex := hourlyTime.Hour()
		if roomHourlyData[roomId][hourIndex] == 0 {
			roomHourlyData[roomId][hourIndex] += float64(sc.StayCount)
		} else {
			roomHourlyData[roomId][hourIndex] += float64(sc.StayCount)
			roomHourlyData[roomId][hourIndex] /= 2
		}

	}

	// 最終的なJSONを作成
	result := make(map[int][]float64)
	for roomId, hourlyData := range roomHourlyData {
		result[roomId] = hourlyData
	}

	for _, sc := range result {
		fmt.Println(sc)
	}

	return result
}
