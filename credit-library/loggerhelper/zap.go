package loggerhelper

import "go.uber.org/zap"

func MakeNewZapProductionLogger() *zap.Logger {
	logger, err := zap.NewProduction()
	defer logger.Sync()
	zap.ReplaceGlobals(logger)

	if err != nil {
		zap.S().Fatalf("configure logger error %v", err)
	}

	return logger
}
