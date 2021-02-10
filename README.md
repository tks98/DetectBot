# DetectBot

DetectBot aims to generalize bot detection across social media platforms. It uses a dataset of ~2800 twitter accounts to train a decision tree classifier.

Depending on the type of social media provided, DetectBot retrieves the necessary features from that user's profile and then makes a prediction using the classifier.

This is mostly a proof of concept for me to learn more about ML, and is not meant to be a serious machine learning project.

## Usage

```bash
go run main.go -c=config.yaml -url=https://twitter.com/username
```

## Config
You will need to apply for a twitter developer account and create a config file. You can pass DetectBot the config using the -c flag.

```
twitter:
  consumerKey: ""
  consumerSecret: ""
  accessToken: ""
  accessSecret: ""
```

## Supported Social Media
* Twitter User Profiles

## TODO
* Twitter Threads
* Facebook/Reddit/Instagram User Profiles and Threads

## License
[GNU GENERAL PUBLIC LICENSE](https://choosealicense.com/licenses/gpl-3.0/)