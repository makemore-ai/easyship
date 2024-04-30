package util

import (
	"github.com/easyship/infra/constant"
	"github.com/easyship/model"
)

const (
	SYSTEM_ERROR_CODE = 50000
	SUCCESS_CODE      = int32(0)
)

func SuccessResp(data interface{}) *model.CommonResponse {
	return &model.CommonResponse{
		Status:     SUCCESS_CODE,
		StatusText: constant.EMPTY_STRING,
		Data:       data,
	}
}

func ErrResp(err error) *model.CommonResponse {
	return &model.CommonResponse{
		Status:     SYSTEM_ERROR_CODE,
		StatusText: err.Error(),
		Data:       constant.EMPTY_STRING,
	}
}
