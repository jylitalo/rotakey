package main

import (
	"testing"

	"github.com/jylitalo/rotakey/cmd"
	"github.com/jylitalo/rotakey/mock"
)

func TestExecute(t *testing.T) {
	execute(func(opts *cmd.Options) {
		opts.AwsCfg = &mock.AwsConfig{}
		opts.FileCfg = &mock.DotAws{}
	})
}
