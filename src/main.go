package main

import (
	"encoding/json"
	"github.com/tks98/DetectBot/config"
	"github.com/tks98/DetectBot/internal/logger"
	"github.com/tks98/DetectBot/pkg/social"
	"github.com/tks98/DetectBot/pkg/twitter"
	"net/url"
	"strings"
)

type BotDetector struct {
	media social.Media
}

func NewBotDetector(media social.Media) BotDetector {
	return BotDetector{media: media}
}

func main() {
	config := config.GetConfig()

	// Init Logger
	loggerConfig, err := json.Marshal(config.Logger)
	if err != nil {
		panic(err)
	}
	logger.InitLogger(loggerConfig)


	for _, url := range config.URL {

		// Get the type of social media depending on the URL
		media := NewSocialMedia(url)

		// Create a new botDetector service with the social media type
		botDetector := NewBotDetector(media)

		// Start the botDetector service
		// Uses the GetFeatures() method of the social media type
		// To retrieve the botdetector needed from the user's profile to run the
		// skllearn script
		err := botDetector.Start()
		if err != nil {
			logger.Log.Errorf(err.Error())
		}

	}
}

func (b BotDetector) Start() error {
	features, err := b.media.GetFeatures() // calls GetFeatures method on the social media type (twitter, facebook, etc)
	if err != nil {
		return err
	}
	bot, confidence,  err := features.RunAIScript()
	if err != nil {
		return err
	}

	if bot {
		logger.Log.Infof("User is %s a bot", confidence)
	} else {
		logger.Log.Infof("User is %s not bot", confidence)
	}
	return nil
}

func NewSocialMedia(url url.URL) social.Media {

	config := config.GetConfig()

	// Depending on the URL, return the correct social media type
	if strings.Contains(url.String(), "https://twitter.com") {
		media := twitter.NewMedia(config.Twitter.ConsumerKey, config.Twitter.ConsumerSecret, config.Twitter.AccessToken, config.Twitter.AccessSecret)
		media.URL = url
		return media
	}

	return nil
}
