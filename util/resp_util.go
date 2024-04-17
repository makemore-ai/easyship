package util

import (
	"github.com/easyship/model"
)

const (
	SYSTEM_ERROR_CODE = 50000
	EMPTY_STR         = ""
	SUCCESS_CODE      = int32(0)
)

func SuccessResp(data interface{}) *model.CommonResponse {
	return &model.CommonResponse{
		Status:     SUCCESS_CODE,
		StatusText: EMPTY_STR,
		Data:       data,
	}
}

func ErrResp(err error) *model.CommonResponse {
	return &model.CommonResponse{
		Status:     SYSTEM_ERROR_CODE,
		StatusText: err.Error(),
		Data:       EMPTY_STR,
	}
}
