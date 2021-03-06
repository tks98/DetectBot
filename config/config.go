package config

import (
	"flag"
	"fmt"
	"github.com/tks98/DetectBot/pkg/twitter"
	"io/ioutil"
	"log"
	"net/url"
	"os"

	"gopkg.in/yaml.v2"
)

var config Config

type Config struct {
	URL          []url.URL
	Socials      []string `yaml:"socials"`
	Twitter struct {
		ConsumerKey    string `yaml:"consumerKey"`
		ConsumerSecret string `yaml:"consumerSecret"`
		AccessToken    string `yaml:"accessToken"`
		AccessSecret   string `yaml:"accessSecret"`
		Media twitter.Media
	} `yaml:"twitter"`
	Logger struct {
		Level            string   `yaml:"level"`
		Encoding         string   `yaml:"encoding"`
		OutputPaths      []string `yaml:"outputPaths"`
		ErrorOutputPaths []string `yaml:"errorOutputPaths"`
		EncoderConfig    struct {
			MessageKey   string `yaml:"messageKey"`
			LevelKey     string `yaml:"levelKey"`
			LevelEncoder string `yaml:"levelEncoder"`
			CallerKey    string `yaml:"callerKey"`
		} `yaml:"encoderConfig"`
	} `yaml:"logger"`

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

	// Check if the URL is valid syntactically
	url, err := url.ParseRequestURI(urlString)
	if err != nil {
		return fmt.Errorf("The URL provided is not a valid URL: %s", url.String())
	}

	config.URL = append(config.URL, *url)

	return nil
}

func GetConfig() *Config {
	return &config
}


// GetConfigs returns the global configs of the controller
func GetConfigs() *Config {
	return &config
}


