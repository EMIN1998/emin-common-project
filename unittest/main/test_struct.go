package main

type UpdateLanguageStatusBatchReq struct {
	ActivityId int64  `json:"activity_id"`
	Language   string `json:"language"`
	Status     int64  `json:"status"`
}

