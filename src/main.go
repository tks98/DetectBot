package main

import (
	"github.com/tks98/Social-Data-Collector/config"
	zap "github.com/tks98/Social-Data-Collector/internal/log"
)

func main() {

	zap.Logger.Info(config.GetConfig())

}
