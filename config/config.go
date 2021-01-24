package config

import (
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

var config Config

type Config struct {
	TwitterCreds struct {
		ConsumerKey    string `yaml:"consumerKey"`
		ConsumerSecret string `yaml:"consumerSecret"`
		AccessToken    string `yaml:"accessToken"`
		AccessSecret   string `yaml:"accessSecret"`
	} `yaml:"twitterCreds"`
}

func init() {
	if len(os.Args) < 2 {
		log.Fatal("No configuration file was specified")
	}

	err := parseConfig(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
}

func parseConfig(configFile string) error {
	file, err := os.Open(configFile)
	if err != nil {
		return err
	}

	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	log.Print(string(bytes))

	err = yaml.UnmarshalStrict(bytes, &config)
	if err != nil {
		return err
	}

	return nil
}

func GetConfig() *Config {
	return &config
}
