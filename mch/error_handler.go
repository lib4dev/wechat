package mch

import (
	"log"
	"os"

	"github.com/micro-plat/hydra/context"
)

type ErrorHandler interface {
	// ServeError 处理回调的错误, 比如 xml 解码出错, return_code != "SUCCESS", result_code != "SUCCESS", ...
	ServeError(*context.Context, error)
}

var DefaultErrorHandler ErrorHandler = ErrorHandlerFunc(defaultErrorHandlerFunc)

type ErrorHandlerFunc func(*context.Context, error)

func (fn ErrorHandlerFunc) ServeError(ctx *context.Context, err error) {
	fn(ctx, err)
}

var errorLogger = log.New(os.Stderr, "[WECHAT_ERROR] ", log.Ldate|log.Ltime|log.Lmicroseconds|log.Llongfile)

func defaultErrorHandlerFunc(ctx *context.Context, err error) {
	errorLogger.Output(3, err.Error())
}
