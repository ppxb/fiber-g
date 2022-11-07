package errorx

import (
	"strings"
	"time"
)

const (
	OK   = 10001
	Fail = 90001
)

const (
	OkMsg = "success"
)

type CodeError struct {
	Code      int    `json:"code"`
	Msg       string `json:"msg"`
	Timestamp int64  `json:"timestamp"`
}

func (c *CodeError) Error() string {
	return c.Msg
}

func (c *CodeError) Data() *CodeErrorResp {
	return &CodeErrorResp{
		Code:      c.Code,
		Msg:       c.Msg,
		Timestamp: c.Timestamp,
	}
}

type CodeErrorResp struct {
	Code      int    `json:"code"`
	Msg       string `json:"msg"`
	Timestamp int64  `json:"timestamp"`
}

func NewCodeError(code int, msg string) error {
	return &CodeError{
		Code:      code,
		Msg:       msg,
		Timestamp: time.Now().Unix(),
	}
}

func NewDefaultError(msg string) error {
	return NewCodeError(Fail, strings.TrimSpace(strings.Split(msg, "=")[2]))
}
