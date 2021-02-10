package botdetector

import (
	"encoding/csv"
	"fmt"
	"github.com/tks98/DetectBot/internal/logger"
	"os"
	"os/exec"
	"path/filepath"
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

// RunAIScript writes the user's features to a temp csv file and then execs the sklearn python script to predict if they are a bot
func (f Features) RunAIScript() (bool, string, error) {

	err := f.writeToCSV()
	if err != nil{
		return  false, "", err
	}


	// Retrieve the absolute path of the python script
	arg, err := filepath.Abs("../pkg/botdetector/bot.py")
	if err != nil {
		return  false, "",  err
	}

	// Remove the nil tag that gets appended to the end of the filepath on macos (for some reason)
	arg = strings.Replace(arg, "<nil>", "", -1)

	// Format the python3 command and pass it and the filepath of the script to the runCommand function to be executed
	command := "python3"
	command = strings.Replace(command, "\\", "", -1 )

	// Execute sklearn script
	output, err := runCommand(command, arg)
	if err != nil {
		return false, "", err
	}


	// Result returned in format "Result-Accuracy"
	// Ex. "BOT-85.948324234234"
	 result := strings.Split(strings.Replace(string(output), "\n", "", -1), "-")
	 if len(result) != 2 {
	 	return false, "" ,fmt.Errorf("unexpected result from AI script: %s", result)
	 }

	 confidence := result[1]

	 // Return result and the confidence level of the prediction
	if result[0] == "BOT" {
		return true, confidence, nil
	}

	return false, confidence, nil
}

// runCommand uses a given command and arguments and executes it
func runCommand(cmdName string, arg string) ([]byte, error) {

	cmd := exec.Command(cmdName, arg)
	logger.Log.Info(cmd.Args)
	out , err := exec.Command(cmdName, arg).Output()
	if err != nil {
		return nil, err
	}

	return out, err
}

// writeToCSV writes the temp user.csv file to be consumed by the sklearn python script
func (f Features) writeToCSV() error {

	// Check for empty features

	if f.Description == "" {
		f.Description = " "
	}
	if f.ScreenName == "" {
		f.ScreenName = " "
	}
	if f.Name == "" {
		f.Name = " "
	}
	if f.Status == "" {
		f.Status = " "
	}

	var data = [][]string{{"screen_name", "name", "description", "status", "verified", "followers_count", "friends_count", "statuses_count", "listedcount", "bot"},
		{f.ScreenName, f.Name, f.Description, f.Status, f.Verified, f.Followers, f.Friends, f.StatusesCount, f.ListedCount, "false"}}

	csvPath, err := filepath.Abs("../pkg/botdetector/user.csv")
	if err != nil {
		return  err
	}

	file, err := os.Create(csvPath)
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


