package xerr

import (
	"fmt"
	"github.com/pkg/errors"
)

/**
常用通用固定错误
*/

type CodeError struct {
	errCode uint32
	errMsg  string
}

func (e *CodeError) Error() string {
	return fmt.Sprintf("ErrCode:%d，ErrMsg:%s", e.errCode, e.errMsg)
}

// GetErrCode Api服务时,返回的错误码
func (e *CodeError) GetErrCode() uint32 {
	return e.errCode
}

// GetErrMsg Api服务时，返回的错误信息
func (e *CodeError) GetErrMsg() string {
	return e.errMsg
}

// NewErr 构造自定义错误
// formatErrString 只会记录在日志里面，不会返回，可以记详细点
func NewErr(errCode uint32, formatErrString string, args ...interface{}) error {
	codeErr := &CodeError{errCode: errCode, errMsg: MapErrMsg(errCode)}
	return errors.Wrapf(codeErr, formatErrString, args)
}
