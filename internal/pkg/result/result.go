package result

import (
	"errors"
)

type Code uint64

var (
	success = newResultError(0, "success")
	fail    = newResultError(1, "fail")
)

var (
	ErrInvalidParams = newResultError(2, "invalid parameters")
)

type Result[T any] struct {
	Code    Code   `json:"code"`
	Message string `json:"message"`
	Data    *T     `json:"data"`
}

type Option[T any] func(result *Result[T])

func WithData[T any](data T) Option[T] {
	return func(result *Result[T]) {
		result.Data = &data
	}
}

func Success[T any](options ...Option[T]) Result[T] {
	r := Result[T]{
		Code:    success.Code,
		Message: success.Message,
		Data:    nil,
	}

	for _, option := range options {
		if option == nil {
			continue
		}
		option(&r)
	}

	return r
}

func Fail[T any](options ...Option[T]) Result[T] {
	r := Result[T]{
		Code:    fail.Code,
		Message: fail.Message,
		Data:    nil,
	}

	for _, option := range options {
		if option == nil {
			continue
		}
		option(&r)
	}

	return r
}

func From[T any](err error, options ...Option[T]) Result[T] {
	if err == nil {
		return Success[T]()
	}

	var ei ErrorInfo
	if !errors.As(err, &ei) {
		return Fail[T]()
	}

	r := Result[T]{
		Code:    ei.Code,
		Message: ei.Message,
		Data:    nil,
	}

	for _, option := range options {
		if option == nil {
			continue
		}
		option(&r)
	}

	return r
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
