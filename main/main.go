package main

import (
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/jylitalo/rotakey/cmd"
)

func execute(opts ...func(*cmd.Options)) error {
	log.SetFormatter(&log.TextFormatter{
		DisableLevelTruncation: true,
		DisableTimestamp:       true,
	})
	return cmd.NewCmd(opts...).Execute()
}

func main() {
	if err := execute(); err == nil {
		os.Exit(0)
	}
	os.Exit(1)
}
