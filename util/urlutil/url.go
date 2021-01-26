package urlutil

import (
	"net/url"
	"strings"
)

func ParseSocial(url *url.URL) string {
	exists := checkIfTwitterUser(url.String())
	if exists {
		return "twitter"
	}

	return ""
}

// checkIfTwitter checks if the url supplied is formatted correctly as a twitter user
// example url: https://twitter.com/userNameGoeshere
func checkIfTwitterUser(url string) bool {

	user := strings.Split(url, "/")

	if len(user) != 4 {
		return false
	}

	if user[0] == "https:" && user[2] == "twitter.com" {
		return true
	}

	return false

}
