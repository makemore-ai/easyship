package util

import (
	"encoding/json"
)

func ToJson(val interface{}) string {
	jsonRes, err := json.Marshal(val)
	if err != nil {
		return ""
	}
	return string(jsonRes)
}
