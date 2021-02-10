package twitter

import (
	"fmt"
	"github.com/dghubble/oauth1"
	"github.com/tks98/Social-Data-Collector/internal/logger"
	"github.com/tks98/Social-Data-Collector/pkg/botdetector"
	"net/url"
	"strings"

	gotwitter "github.com/dghubble/go-twitter/twitter"
)

type Media struct {
	Client *gotwitter.Client
	URL    url.URL
}

func NewMedia(consumerKey, consumerSecret, accessToken, accessSecret string) *Media {

	// Setup oauth client
	oauthConfig := oauth1.NewConfig(consumerKey, consumerSecret)
	oauthToken := oauth1.NewToken(accessToken, accessSecret)
	httpClient := oauthConfig.Client(oauth1.NoContext, oauthToken)

	// Create twitter client
	twitterClient := gotwitter.NewClient(httpClient)

	twitter := &Media{
		Client: twitterClient,
	}

	return twitter
}

// GetFeatures uses the twitter API to visit a user profile and retrieve the fields needed to populate the features type
func (t Media) GetFeatures() (botdetector.Features, error) {

	// Parse the username from the URL
	var username []string
	splitURL := strings.Split(t.URL.String(), "/")

	if len(splitURL) == 4 {
		username = append(username, splitURL[3])
	} else {
		return botdetector.Features{}, fmt.Errorf("unsupported twitter url, please supply it in the format: https://twitter.com/username")
	}

	logger.Log.Infof("Username: %s", username)

	// Use client to retrieve user features
	users, resp, err := t.Client.Users.Lookup(&gotwitter.UserLookupParams{
		ScreenName: username,
	})

	if err != nil {
		logger.Log.Errorf(err.Error())
		logger.Log.Info(resp.StatusCode)
		return botdetector.Features{}, err
	}

	if len(users) > 1 {
		return botdetector.Features{}, fmt.Errorf("more than one user was retured with username %s", username)
	}

	user := users[0]

	var status string


	features := botdetector.Features{
		ScreenName:    user.ScreenName,
		Name:          user.Name,
		Description:   user.Description,
		Status:        status,
		Verified:      fmt.Sprintf("%v", user.Verified),
		Followers:     fmt.Sprintf("%v", user.FollowersCount),
		Friends:       fmt.Sprintf("%v", user.FriendsCount),
		StatusesCount: fmt.Sprintf("%v", user.StatusesCount),
		ListedCount:   fmt.Sprintf("%v", user.ListedCount),
		Bot:           fmt.Sprintf("%v", false),
	}

	return features, nil
}
