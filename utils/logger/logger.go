package logger

import (
	"context"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func GetLoggerFromContext(ctx context.Context) *zerolog.Logger {
	if l, ok := ctx.Value(keyLogger).(*zerolog.Logger); ok {
		return l
	}

	// default logger
	return &log.Logger
}

func AddLoggerToContext(l *zerolog.Logger, ctx context.Context) context.Context {
	return context.WithValue(ctx, keyLogger, l)
}
