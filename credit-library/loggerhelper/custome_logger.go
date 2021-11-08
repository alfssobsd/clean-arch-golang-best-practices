package loggerhelper

import (
	"context"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
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

func (l *CustomLogger) InfoWithTracing(ctx context.Context, msg string, fields ...zapcore.Field) {
	//append fields
	l.zapLogger.Info(msg, fields...)
}
func (l *CustomLogger) InfoNoTracing(msg string, fields ...zapcore.Field) {
	l.zapLogger.Info(msg, fields...)
}

func (l *CustomLogger) InfofWithTracing(ctx context.Context, template string, args ...interface{}) {
	l.zapLogger.Sugar().With(ctx).Infof(template, args)
}

func (l *CustomLogger) InfofNoTracing(template string, args ...interface{}) {
	l.zapLogger.Sugar().Infof(template, args)
}

func (l *CustomLogger) ErrorWithTracing(ctx context.Context, msg string, fields ...zapcore.Field) {
	//append fields
	l.zapLogger.Error(msg, fields...)
}
func (l *CustomLogger) ErrorNoTracing(msg string, fields ...zapcore.Field) {
	l.zapLogger.Error(msg, fields...)
}

func (l *CustomLogger) ErrorfWithTracing(ctx context.Context, template string, args ...interface{}) {
	l.zapLogger.Sugar().With().Errorf(template, args)
}

func (l *CustomLogger) ErrorfNoTracing(template string, args ...interface{}) {
	l.zapLogger.Sugar().Errorf(template, args)
}
func (l *CustomLogger) WarnWithTracing(ctx context.Context, msg string, fields ...zapcore.Field) {
	//append fields
	l.zapLogger.Warn(msg, fields...)
}
func (l *CustomLogger) WarnNoTracing(msg string, fields ...zapcore.Field) {
	l.zapLogger.Warn(msg, fields...)
}

func (l *CustomLogger) WarnfWithTracing(ctx context.Context, template string, args ...interface{}) {
	l.zapLogger.Sugar().With().Warnf(template, args)
}

func (l *CustomLogger) WarnfNoTracing(template string, args ...interface{}) {
	l.zapLogger.Sugar().Warnf(template, args)
}

func (l *CustomLogger) DebugfWithTracing(ctx context.Context, template string, args ...interface{}) {
	l.zapLogger.Sugar().With().Debugf(template, args)
}

func (l *CustomLogger) DebugfNoTracing(template string, args ...interface{}) {
	l.zapLogger.Sugar().Debugf(template, args)
}

func (l *CustomLogger) PanicfWithTracing(ctx context.Context, template string, args ...interface{}) {
	l.zapLogger.Sugar().With().Panicf(template, args)
}

func (l *CustomLogger) PanicfNoTracing(template string, args ...interface{}) {
	l.zapLogger.Sugar().Panicf(template, args)
}

func (l *CustomLogger) FatalfWithTracing(template string, args ...interface{}) {
	l.zapLogger.Sugar().With().Panicf(template, args)
}

func (l *CustomLogger) FatalfNoTracing(template string, args ...interface{}) {
	l.zapLogger.Sugar().Panicf(template, args)
}
