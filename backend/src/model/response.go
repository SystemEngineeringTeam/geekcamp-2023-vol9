package model

import "time"

type GetStayCountResponseModel struct {
	Buildings []GetStayCountBuildingModel			`json:"building"`
}

type GetCongestionResponseModel struct {
	Building []GetCongestionBuildingModel			`json:"building"`
}

type PostStayCountRequestModel struct {
	Time time.Time 				`json:"time"`
	Headcount int 				`json:"headcount"`
}

type GetStayCountHistoryRequestModel struct {
	Histories []GetStayCountHistoryModel	`json:"histories"`
}

type GetStayCountHistoryModel struct {
	StayCount [24]int	`json:"0"`
}