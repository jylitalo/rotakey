package cmd

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/jylitalo/rotakey"
	"github.com/jylitalo/rotakey/internal"
	"github.com/jylitalo/rotakey/types"
)

type Options struct {
	Use     string
	AwsCfg  types.AwsConfig
	FileCfg types.DotAws
	Rotate  types.Rotate
}

func NewCmd(optFns ...func(*Options)) *cobra.Command {
	opts := &Options{
		Use:     "rotakey",
		AwsCfg:  &rotakey.AwsConfig{},
		FileCfg: &rotakey.DotAws{},
		Rotate:  &rotakey.Rotate{},
	}
	for _, fn := range optFns {
		fn(opts)
	}
	cmd := &cobra.Command{
		Use:   opts.Use,
		Short: "Rotate AWS IAM credentials",
		Long: `Rotate AWS IAM credentials by
1. creating new access keys into IAM
2. updating new access keys into ~/.aws/credentials
3. remove current access keys from IAM
`,
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			flags := cmd.Flags()
			debug, errD := flags.GetBool("debug")
			verbose, errV := flags.GetBool("verbose")
			if err := internal.CoalesceError(errD, errV); err != nil {
				return err
			}
			switch {
			case debug:
				log.SetLevel(log.DebugLevel)
			case verbose:
				log.SetLevel(log.InfoLevel)
			default:
				log.SetLevel(log.WarnLevel)
			}
			err := opts.Rotate.Execute(opts.AwsCfg, opts.FileCfg)
			if err != nil {
				return err
			}
			fmt.Println("Task completed.")
			return nil
		},
	}
	cmd.Flags().Bool("debug", false, "Debug output")
	cmd.Flags().BoolP("verbose", "v", false, "Verbose output")
	return cmd
}
