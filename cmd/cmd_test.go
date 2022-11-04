package cmd

import (
	"testing"

	"github.com/jylitalo/rotakey/mock"
)

func TestNoFlags(t *testing.T) {
	NewCmd(func(o *Options) {
		o.Rotate = &mock.Rotate{}
	}).Execute()
}

func TestDebugFlag(t *testing.T) {
	cmd := NewCmd(func(o *Options) { o.Rotate = &mock.Rotate{} })
	cmd.SetArgs([]string{"--debug"})
	cmd.Execute()
}

func TestVerboseFlag(t *testing.T) {
	cmd := NewCmd(func(o *Options) { o.Rotate = &mock.Rotate{} })
	cmd.SetArgs([]string{"--verbose"})
	cmd.Execute()
}
