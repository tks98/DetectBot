package main

import (
	"github.com/tks98/Social-Data-Collector/config"
	"github.com/tks98/Social-Data-Collector/internal/log"
	"github.com/tks98/Social-Data-Collector/util/strutil"
)

func main() {

	// Parse the URL given from the user to determine the type of social media it is for
	socialType, err := strutil.ParseURL(config.GetConfig().URL)
	if err != nil {
		log.Print.Fatal(err)
	}

	if socialType == "twitter" {
		log.Print.Info("It works!")
	} else {
		log.Print.Info("Lame")
	}

}
