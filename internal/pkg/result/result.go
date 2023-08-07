package result

import (
	"github.com/pkg/errors"
)

type Code uint64

var (
	success = newResultError(0, "success")
	fail    = newResultError(1, "fail")
)

type Result[T any] struct {
	Code    Code   `json:"code"`
	Message string `json:"message"`
	Data    T      `json:"data"`
}

func Success(data any) Result[any] {
	return Result[any]{
		Code:    success.Code,
		Message: success.Message,
		Data:    data,
	}
}

func Fail(data any) Result[any] {
	return Result[any]{
		Code:    fail.Code,
		Message: fail.Message,
		Data:    data,
	}
}

func From(err error, data any) Result[any] {
	if err == nil {
		return Success(nil)
	}

	var ei ErrorInfo
	if !errors.As(err, &ei) {
		return Fail(nil)
	}

	return Result[any]{
		Code:    ei.Code,
		Message: ei.Message,
		Data:    data,
	}
}

type ErrorInfo struct {
	Code    Code   `json:"code"`
	Message string `json:"message"`
}

func newResultError(code Code, message string) ErrorInfo {
	return ErrorInfo{
		Code:    code,
		Message: message,
	}
}

func (e ErrorInfo) Error() string {
	return e.Message
}

func (e ErrorInfo) Is(target error) bool {
	return e.Message == target.Error()
}
