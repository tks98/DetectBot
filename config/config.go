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
	User         string
	Thread       string
	TwitterCreds struct {
		ConsumerKey    string `yaml:"consumerKey"`
		ConsumerSecret string `yaml:"consumerSecret"`
		AccessToken    string `yaml:"accessToken"`
		AccessSecret   string `yaml:"accessSecret"`
	} `yaml:"twitterCreds"`
}

func init() {
	config := flag.String("c", "", "Specifies path to config file")
	user := flag.String("u", "", "Specifies a URL to a user to check")
	thread := flag.String("t", "", "Specifies a thread to check")
	flag.Parse()

	if *config == "" {
		log.Fatal("You need to specify a configuration file")
	}

	err := parseConfig(*config, *user, *thread)
	if err != nil {
		log.Fatal(err)
	}
}

func parseConfig(configFile string, user string, thread string) error {
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

	if user == "" && thread == "" {
		return fmt.Errorf("You need to specify a link to a user, or a thread. Use the -help flag for assistance")
	} else if user == "" {
		config.Thread = thread
	} else if thread == "" {
		config.User = user
	} else {
		return fmt.Errorf("You specified too many flags!! Type -help for help")
	}

	return nil
}

func GetConfig() *Config {
	return &config
}
