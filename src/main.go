package main

import (
	"github.com/tks98/Social-Data-Collector/config"
	i "github.com/tks98/Social-Data-Collector/internal"
)

func main() {
	i.Logger.Info(config.GetConfig())
}
