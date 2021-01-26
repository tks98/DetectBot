package social

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/tks98/Social-Data-Collector/pkg/twitter"
	"github.com/tks98/Social-Data-Collector/util/strutil"
)

type User struct {
	Username string
	Exists   bool
	URL      url.URL
}

type Post struct {
	ID     string
	Exists bool
	URL    url.URL
}

type Kind struct {
	Website  string
	LinkType string
}

type SocialMedia struct {
	User User
	Post Post
	Kind Kind
}

// CheckValidity determines if a social media link points to a valid user or post
func CheckValidity(url url.URL, socials []string) (bool, error) {

	// Determine the type of the URL
	socialMedia := ParseSocial(url)

	// Check if the type is supported
	if !strutil.Contains(socials, socialMedia.Kind.Website) {
		return false, fmt.Errorf("Social media website is not supported: %s", socialMedia.Kind.Website)
	}

	// Check if the user or post exists
	switch socialMedia.Kind.LinkType {
	case "twitterUser":
		return twitter.GetClient().CheckIfUserExists(socialMedia.User.Username)
	}

	return false, nil

}

// ParseSocial takes a URL and parses it to determine the type of social media
func ParseSocial(url url.URL) SocialMedia {

	var socialMedia SocialMedia
	exists, username := checkIfTwitterUser(url.String())
	if exists {
		user := User{
			Username: username,
			Exists:   false,
			URL:      url,
		}
		socialMedia.User = user
		socialMedia.Kind = Kind{
			Website:  "twitter",
			LinkType: "twitterUser",
		}
		return socialMedia
	}

	return socialMedia
}

// checkIfTwitter checks if the url supplied is formatted correctly as a twitter user
// example url: https://twitter.com/userNameGoeshere
func checkIfTwitterUser(url string) (bool, string) {

	user := strings.Split(url, "/")

	if len(user) != 4 {
		return false, ""
	}

	if user[0] == "https:" && user[2] == "twitter.com" {
		return true, user[3]
	}

	return false, ""

}
