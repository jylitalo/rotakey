package main

import (
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/jylitalo/rotakey/cmd"
)

func execute(params cmd.NewCmdInput) error {
	log.SetFormatter(&log.TextFormatter{
		DisableLevelTruncation: true,
		DisableTimestamp:       true,
	})
	return cmd.NewCmd(params).Execute()
}

func main() {
	if err := execute(cmd.NewCmdInput{}); err == nil {
		os.Exit(0)
	}
	os.Exit(1)
}
