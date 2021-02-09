package botdetector

// Features is a type that will get populated before being sent to the skllearn script for processing
type Features struct {
	ScreenName    string
	Name          string
	Description   string
	Status        string
	Verified      bool
	Followers     int
	Friends       int
	StatusesCount int
	ListedCount   int
	Bot           bool
}

func (f Features) RunAIScript() error {
	return nil
}
