package main

import (
	"github.com/tks98/Social-Data-Collector/config"
	"github.com/tks98/Social-Data-Collector/internal/log"
)

func main() {
	log.Print.Info(config.GetConfig().SocialType)
}
