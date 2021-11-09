package loggerhelper

import (
	"context"
	"go.uber.org/zap"
)

type CustomLogger struct {
	zapLogger *zap.Logger
}

func NewCustomLogger(options ...zap.Option) *CustomLogger {
	zapLogger, err := zap.NewProduction(options...)
	defer zapLogger.Sync()
	zap.ReplaceGlobals(zapLogger)

	if err != nil {
		zap.S().Fatalf("configure logger error %v", err)
	}

	return &CustomLogger{zapLogger: zapLogger}
}
func (l *CustomLogger) WithTracing(ctx context.Context) *zap.Logger {
	return l.zapLogger.With()
}

func (l *CustomLogger) NoTracing() *zap.Logger {
	return l.zapLogger.With()
}

func (l *CustomLogger) SugarWithTracing(ctx context.Context) *zap.SugaredLogger {
	return l.zapLogger.With().Sugar()
}

func (l *CustomLogger) SugarNoTracing() *zap.SugaredLogger {
	return l.zapLogger.With().Sugar()
}
