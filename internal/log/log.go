package log

import "go.uber.org/zap"

// Logger is a shared global logger
var Print *zap.SugaredLogger

func init() {
	// setup zap logger according to their quick start guide
	logger, err := zap.NewProduction()
	if err != nil {
		logger.Fatal(err.Error())
	}
	defer logger.Sync()
	sugar := logger.Sugar()
	Print = sugar
}
