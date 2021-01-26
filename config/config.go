package config

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
	"os"

	"github.com/tks98/Social-Data-Collector/util/strutil"
	"github.com/tks98/Social-Data-Collector/util/urlutil"
	"gopkg.in/yaml.v2"
)

var config Config

type Config struct {
	URL          url.URL
	SocialType   string
	Socials      []string `yaml:"socials"`
	TwitterCreds struct {
		ConsumerKey    string `yaml:"consumerKey"`
		ConsumerSecret string `yaml:"consumerSecret"`
		AccessToken    string `yaml:"accessToken"`
		AccessSecret   string `yaml:"accessSecret"`
	} `yaml:"twitterCreds"`
}

func init() {
	config := flag.String("c", "", "Specifies path to config file")
	url := flag.String("url", "", "Specifies a URL")
	flag.Parse()

	if *config == "" {
		log.Fatal("You need to specify a configuration file")
	}

	err := parseConfig(*config, *url)
	if err != nil {
		log.Fatal(err)
	}
}

func parseConfig(configFile string, urlString string) error {
	file, err := os.Open(configFile)
	if err != nil {
		return err
	}

	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	err = yaml.UnmarshalStrict(bytes, &config)
	if err != nil {
		return err
	}

	if urlString == "" {
		return fmt.Errorf("You need to specify a URL of a user or a post to check")
	}

	// Check if the URL is valid
	url, err := url.ParseRequestURI(urlString)
	if err != nil {
		return fmt.Errorf("The URL provided is not valid. URL: %s", url.String())
	}

	// Determine the type of the URL
	socialType := urlutil.ParseSocial(url)

	// Check if the type is supported
	if !strutil.Contains(config.Socials, socialType) {
		return fmt.Errorf(err.Error())
	}

	config.URL = *url
	config.SocialType = socialType
	return nil
}

func GetConfig() *Config {
	return &config
}
