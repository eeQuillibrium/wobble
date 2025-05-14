package logger

import (
	"context"
	"go.uber.org/zap"
)

type logger *zap.SugaredLogger

type ctxKey struct{}

var log logger

func init() {
	log = zap.NewExample().Sugar()
}

func GetLogger() logger {
	return log
}

func Ctx(ctx context.Context) *zap.SugaredLogger {
	if l, ok := ctx.Value(ctxKey{}).(*zap.SugaredLogger); ok {
		return l
	}
	return log
}
