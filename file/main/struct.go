package main

type ReadData struct {
	Language   string `json:"language"`
	ActivityId int64  `json:"activity_id"`
	Status     int64  `json:"status"`
}

type ReadDataList []*ReadData

type saveData struct {
	ActivityId int64  `json:"activity_id"`
	Reason     string `json:"err"`
}

type MfoKeyWord struct {
	KeywordID int64  `json:"keyword_id"`
	CountryId int64  `json:"country_id"`
	CityID    int64  `json:"city_id"`
	Language  string `json:"language"`
}
