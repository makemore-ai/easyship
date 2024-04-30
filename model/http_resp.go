package model

type CommonResponse struct {
	Status     int32       `json:"status"`
	StatusText string      `json:"status_text"`
	Data       interface{} `json:"data"`
}

type StreamResult struct {
	Id    int32
	Event string
	Data  *string
}

type IdCal struct {
	NowId *int32
}
