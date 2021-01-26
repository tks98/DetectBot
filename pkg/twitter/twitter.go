package twitter

import (
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/tks98/Social-Data-Collector/config"
)

var Client *twitter.Client

func init() {
	// Get pointer to application config for twitter creds
	config := config.GetConfig()
	// Setup oauth client
	oauthConfig := oauth1.NewConfig(config.TwitterCreds.ConsumerKey, config.TwitterCreds.ConsumerSecret)
	oauthToken := oauth1.NewToken(config.TwitterCreds.AccessToken, config.TwitterCreds.AccessSecret)
	httpClient := oauthConfig.Client(oauth1.NoContext, oauthToken)

	// Create twitter client
	twitterClient := twitter.NewClient(httpClient)
	Client = twitterClient

}
