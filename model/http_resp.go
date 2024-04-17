package model

type CommonResponse struct {
	Status     int32       `json:"status"`
	StatusText string      `json:"status_text"`
	Data       interface{} `json:"data"`
}
