package mch

import (
	"log"
	"os"

	"github.com/micro-plat/hydra"
)

type ErrorHandler interface {
	// ServeError 处理回调的错误, 比如 xml 解码出错, return_code != "SUCCESS", result_code != "SUCCESS", ...
	ServeError(hydra.IContext, error)
}

var DefaultErrorHandler ErrorHandler = ErrorHandlerFunc(defaultErrorHandlerFunc)

type ErrorHandlerFunc func(hydra.IContext, error)

func (fn ErrorHandlerFunc) ServeError(ctx hydra.IContext, err error) {
	fn(ctx, err)
}

var errorLogger = log.New(os.Stderr, "[WECHAT_ERROR] ", log.Ldate|log.Ltime|log.Lmicroseconds|log.Llongfile)

func defaultErrorHandlerFunc(ctx hydra.IContext, err error) {
	errorLogger.Output(3, err.Error())
}
