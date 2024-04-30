package util

import (
	"encoding/json"
	"io"
)

func ToJson(val interface{}) string {
	jsonRes, err := json.Marshal(val)
	if err != nil {
		return ""
	}
	return string(jsonRes)
}

func ParseJson(val string) (map[string]interface{}, error) {
	res := map[string]interface{}{}
	err := json.Unmarshal([]byte(val), &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func DecodeJson(reader io.Reader) (map[string]interface{}, error) {
	decoder := json.NewDecoder(reader)
	res := map[string]interface{}{}
	if err := decoder.Decode(&res); err != nil {
		// no log交给上层
		return res, err
	}
	return res, nil
}
