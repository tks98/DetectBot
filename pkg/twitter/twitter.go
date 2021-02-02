package twitter

import (

	"github.com/tks98/Social-Data-Collector/pkg/botdetector"

	gotwitter "github.com/dghubble/go-twitter/twitter"

)


type Media struct {
	Client *gotwitter.Client
}

// GetFeatures uses the twitter API to visit a user profile and retrieve the fields needed to populate the features type
func (t Media) GetFeatures() botdetector.Features {
	return botdetector.Features{}
}
