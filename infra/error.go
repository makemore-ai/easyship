package infra

import (
	"fmt"
	"github.com/easyship/util"
)

const DEFAULT_ERR_MSG = "系统错误"

// 系统异常
type SystemError struct {
	errCode int
	errMsg  string
}

func (err *SystemError) Error() string {
	return fmt.Sprintf("SystemError err code:%v, errMsg:%v", err.errCode, err.errMsg)
}
func NewDefaultSystemError() error {
	return &SystemError{
		errCode: util.SYSTEM_ERROR_CODE,
		errMsg:  DEFAULT_ERR_MSG,
	}
}
func NewSystemError(errMsg string) error {
	return &SystemError{
		errCode: util.SYSTEM_ERROR_CODE,
		errMsg:  errMsg,
	}
}
