package features

// Features is a type that will get populated before being sent to the skllearn script for processing
type Features struct {
	ScreenName    string
	Name          string
	Description   string
	Status        string
	Verified      string
	Followers     int
	Friends       int
	StatusesCount int
	ListedCount   int
	bot           bool
}
