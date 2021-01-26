package config

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

var config Config

type Config struct {
	URL          string
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

func parseConfig(configFile string, url string) error {
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

	if url == "" {
		return fmt.Errorf("You need to specify a URL of a user or a post to check")
	}

	config.URL = url

	return nil
}

func GetConfig() *Config {
	return &config
}
