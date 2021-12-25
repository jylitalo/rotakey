package cmd

import (
	"testing"

	"github.com/jylitalo/rotakey"
)

func TestNoFlags(t *testing.T) {
	NewCmd(NewCmdInput{NewExec: rotakey.NewMockExec}).Execute()
}

func TestDebugFlag(t *testing.T) {
	cmd := NewCmd(NewCmdInput{NewExec: rotakey.NewMockExec})
	cmd.SetArgs([]string{"--debug"})
	cmd.Execute()
}

func TestVerboseFlag(t *testing.T) {
	cmd := NewCmd(NewCmdInput{NewExec: rotakey.NewMockExec})
	cmd.SetArgs([]string{"--verbose"})
	cmd.Execute()
}
