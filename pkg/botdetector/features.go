package botdetector

import (
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
	Verified      bool
	Followers     int
	Friends       int
	StatusesCount int
	ListedCount   int
	Bot           bool
}

func (f Features) RunAIScript() error {

	logger.Log.Info("run ai script")
	arg := "/Users/tsmith/Documents/Projects/Social-Data-Collector/pkg/botdetector/bot.py"
	command := "python3"
	command = strings.Replace(command, "\\", "", -1 )
	err := runCommand(command, arg)
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
