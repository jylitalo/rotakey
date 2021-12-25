package main

import (
	log "github.com/sirupsen/logrus"

	"github.com/jylitalo/rotakey/cmd"
)

func execute(params cmd.NewCmdInput) {
	log.SetFormatter(&log.TextFormatter{
		DisableLevelTruncation: true,
		DisableTimestamp:       true,
	})
	cmd.NewCmd(params).Execute()

}
func main() {
	execute(cmd.NewCmdInput{})
}
