package util

import (
	"bytes"
	"io"
)

const (
	CONTENT_TYPE     = "Content-Type"
	APPLICATION_JSON = "application/json"
)

func ParseBody(value interface{}) io.Reader {
	jsonResult := ToJson(value)
	return bytes.NewReader([]byte(jsonResult))
}
