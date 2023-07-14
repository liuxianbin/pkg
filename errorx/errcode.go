package errorx

import (
	"fmt"
)

type Error struct {
	httpCode int // httpCode
	code     int
	message  string
	details  []string
}

func (e *Error) HttpCode() int {
	return e.httpCode
}

func (e *Error) Code() int {
	return e.code
}

func (e *Error) Message() string {
	return e.message
}

func (e *Error) Details() []string {
	return e.details
}

var codes = map[int]string{}

func NewError(httpCode, code int, message string) *Error {
	if _, ok := codes[code]; ok {
		panic(fmt.Sprintf("错误码 %d 已存在", code))
	}
	codes[code] = message
	return &Error{httpCode: httpCode, code: code, message: message}
}

func (e *Error) Error() string {
	return fmt.Sprintf("错误码: %d, 错误信息: %s", e.code, e.message)
}

func (e *Error) WithDetails(details ...string) *Error {
	_e := *e
	_e.details = []string{}
	for _, d := range details {
		_e.details = append(_e.details, d)
	}
	return &_e
}
