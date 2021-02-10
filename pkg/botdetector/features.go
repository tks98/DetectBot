package botdetector

import (
	"encoding/csv"
	"github.com/tks98/Social-Data-Collector/internal/logger"
	"os"
	"os/exec"
	"strings"
)

// Features is a type that will get populated before being sent to the skllearn script for processing
type Features struct {
	ScreenName    string
	Name          string
	Description   string
	Status        string
	Verified      string
	Followers     string
	Friends       string
	StatusesCount string
	ListedCount   string
	Bot           string
}

func (f Features) RunAIScript() error {

	err := f.writeToCSV()
	if err != nil{
		return err
	}

	logger.Log.Info("run ai script")
	arg := "/Users/tsmith/Documents/Projects/Social-Data-Collector/pkg/botdetector/bot.py"
	command := "python3"
	command = strings.Replace(command, "\\", "", -1 )
	err = runCommand(command, arg)
	if err != nil {
		return err
	}

	return nil
}

func runCommand(cmdName string, arg ...string) error {
	cmd := exec.Command(cmdName, arg...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}


func (f Features) writeToCSV() error {
	var data = [][]string{{"screen_name", "name", "description", "status", "verified", "followers_count", "friends_count", "statuses_count", "listedcount", "bot"},
		{f.ScreenName, f.Name, f.Description, f.Status, f.Verified, f.Followers, f.Friends, f.StatusesCount, f.ListedCount, "false"}}


	file, err := os.Create("user.csv")
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, value := range data {
		err := writer.Write(value)
		if err != nil {
			return err
		}
	}

	return nil
}


