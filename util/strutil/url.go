package strutil

import (
	"fmt"

	"github.com/tks98/Social-Data-Collector/internal/log"

	"net/url"
	"strings"

	"github.com/tks98/Social-Data-Collector/pkg/twitter"
)

// ParseURL determines what type of social media platform and user/thread was specified by the user
func ParseURL(url *url.URL) (string, error) {
	exists, err := checkIfTwitterUser(url.String())
	if err != nil {
		return "", err
	}
	if exists {
		return "twitter", nil
	}
	return "", nil
}

// checkIfTwitter checks if the url supplied is that of a twitter user
// example url: https://twitter.com/userNameGoeshere
func checkIfTwitterUser(url string) (bool, error) {

	user := strings.Split(url, "/")
	log.Print.Info(user)

	if len(user) != 4 {
		return false, nil
	}

	if user[0] == "https:" && user[2] == "twitter.com" {
		exists, err := twitter.GetClient().CheckIfUserExists(user[3])
		if err != nil {
			return false, err
		}

		if !exists {
			return false, nil
		}
	}
	fmt.Println(user)
	return true, nil
}
