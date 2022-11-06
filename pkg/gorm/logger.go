package gorm

import (
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm/logger"
)

type Logger struct {
	logger.Writer
}

func (l Logger) Printf(msg string, data ...interface{}) {
	logx.Errorf(msg, data)
}
