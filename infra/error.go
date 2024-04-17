package infra

import (
	"fmt"
	"github.com/easyship/util"
)

// 系统异常
type SystemError struct {
	errCode int
	errMsg  string
}

func (err *SystemError) Error() string {
	return fmt.Sprintf("SystemError err code:%v, errMsg:%v", err.errCode, err.errMsg)
}

func NewSystemError(errMsg string) error {
	return &SystemError{
		errCode: util.SYSTEM_ERROR_CODE,
		errMsg:  errMsg,
	}
}
