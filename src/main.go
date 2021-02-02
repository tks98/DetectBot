package main

import (
	gotwitter "github.com/dghubble/go-twitter/twitter"
	"github.com/tks98/Social-Data-Collector/config"
	"github.com/tks98/Social-Data-Collector/pkg/social"
	"github.com/tks98/Social-Data-Collector/pkg/twitter"
	"net/url"
)




type BotDetector struct {
	media social.Media
}

func NewBotDetector(media social.Media) BotDetector {
	return BotDetector{media: media}
}



func main() {
	config := config.GetConfig()

	// Get the type of social media depending on the URL
	media := NewSocialMedia(config.URL[0])

	// Create a new botDetector service with the social media type
	botDetector := NewBotDetector(media)

	// Start the botDetector service
	// Uses the GetFeatures() method of the social media type
	// To retrieve the botdetector needed from the user's profile to run the
	// skllearn script
	botDetector.Start()
}
func (b BotDetector) Start() bool {
	features := b.media.GetFeatures() // calls GetFeatures method on the socal media type (twitter, facebook, etc)
	features.RunAIScript()
	return true
}


func NewSocialMedia(url url.URL) social.Media {

	// Depending on the URL, return the correct social media type
	if url.String() == "twitter.com" {
		return twitter.Media{Client: &gotwitter.Client{}}
	}

	return nil
}





