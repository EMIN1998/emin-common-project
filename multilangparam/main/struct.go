package main


type Data struct {
	ActivityId int64 `json:"activity_id"`
	Status int64 `json:"status"`
	Language string `json:"language"`
}

type RequestParam struct {
	Language string `json:"language"`
	ActivityIds  string `json:"activity_ids"`
}
