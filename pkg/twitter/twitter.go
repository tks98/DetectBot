package twitter

import (
	"fmt"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/tks98/Social-Data-Collector/config"
	"github.com/tks98/Social-Data-Collector/internal/log"
)

type Client struct {
	HTTPClient *twitter.Client
}

var client *Client

func init() {
	// Get pointer to application config for twitter creds
	config := config.GetConfig()
	// Setup oauth client
	oauthConfig := oauth1.NewConfig(config.TwitterCreds.ConsumerKey, config.TwitterCreds.ConsumerSecret)
	oauthToken := oauth1.NewToken(config.TwitterCreds.AccessToken, config.TwitterCreds.AccessSecret)
	httpClient := oauthConfig.Client(oauth1.NoContext, oauthToken)

	// Create twitter client
	twitterClient := twitter.NewClient(httpClient)

	twitter := &Client{
		HTTPClient: twitterClient,
	}

	client = twitter

}

func GetClient() *Client {
	return client
}

func (c *Client) CheckIfUserExists(username string) (bool, error) {
	usernameSlice := []string{username}
	log.Print.Info(username)
	user, resp, err := c.HTTPClient.Users.Lookup(&twitter.UserLookupParams{
		ScreenName: usernameSlice,
	})
	if err != nil {
		return false, err
	}

	if resp.StatusCode != 200 {
		if resp.StatusCode == 404 {
			return false, fmt.Errorf("User %s does not exist", username)
		} else {
			return false, fmt.Errorf("Request to get user %s was not successful. Status code: %d", username, resp.StatusCode)
		}
	}

	log.Print.Info(user)
	return true, err

}
