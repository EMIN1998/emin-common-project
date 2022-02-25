package main

type UpdateBatchReq struct {
	ActivityIds  string `json:"activity_ids"`
	Language     string `json:"language"`
	TargetStatus int    `json:"target_status"` // 如果是机翻发布则为5，如果是美国英文发布则是1
}

type UpdateLanguageStatusBatchReq struct {
	ActivityId int64  `json:"activity_id"`
	Language   string `json:"language"`
	Status     int64  `json:"status"`
}
