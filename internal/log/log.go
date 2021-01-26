package internal

import "go.uber.org/zap"

// Logger is a shared global logger
var Logger *zap.SugaredLogger

func init() {
	// setup zap logger according to their quick start guide
	logger, err := zap.NewProduction()
	if err != nil {
		logger.Fatal(err.Error())
	}
	defer logger.Sync()
	sugar := logger.Sugar()
	Logger = sugar
}
