package logger

import (
	"fmt"
	
	"go.uber.org/zap"
)

type Writer struct{
}

func (w Writer) Printf(format string,args ...interface{}) {
    // log.Infof(format, args...)
	zlogger := zap.L()
	zlogger.Info(fmt.Sprintf(format, args...))
}