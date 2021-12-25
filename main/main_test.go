package main

import (
	"testing"

	"github.com/jylitalo/rotakey"
	"github.com/jylitalo/rotakey/cmd"
)

func TestExecute(t *testing.T) {
	execute(cmd.NewCmdInput{NewExec: rotakey.NewMockExec})
}
