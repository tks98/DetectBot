package main

import (
	"github.com/tks98/Social-Data-Collector/config"
	"github.com/tks98/Social-Data-Collector/internal/log"
	"github.com/tks98/Social-Data-Collector/internal/social"
)

func main() {
	config := config.GetConfig()

	// Depending on the social media type provided, determine if the link is to a valid post, or user
	valid, err := social.CheckValidity(config.URL, config.Socials)
	if err != nil {
		log.Print.Fatal(err)
	}

	if valid {
		log.Print.Info("Link is valid")
		return
	} else {
		log.Print.Info("Link is not valid")
	}

}
