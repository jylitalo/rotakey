package main

import (
	log "github.com/sirupsen/logrus"

	"github.com/jylitalo/rotakey"
	"github.com/jylitalo/rotakey/cmd"
)

func main() {
	log.SetFormatter(&log.TextFormatter{
		DisableLevelTruncation: true,
		DisableTimestamp:       true,
	})
	cmd.NewRotakeyCmd(rotakey.NewAwsConfig, rotakey.NewDotAws).Execute()
}
