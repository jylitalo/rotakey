package cmd

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/jylitalo/rotakey"
)

func NewRotakeyCmd(newAwsConfig func() (rotakey.AwsConfigIface, error), newDotAws func() (rotakey.DotAwsIface, error)) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "rotakey",
		Short: "Rotate AWS IAM credentials",
		Long: `Rotate AWS IAM credentials by
1. creating new access keys into IAM
2. updating new access keys into ~/.aws/credentials
3. remove current access keys from IAM
`,
		RunE: func(cmd *cobra.Command, args []string) error {
			flags := cmd.Flags()
			debug, errD := flags.GetBool("debug")
			verbose, errV := flags.GetBool("verbose")
			if err := rotakey.CoalesceError(errD, errV); err != nil {
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
			if err := rotakey.Execute(newAwsConfig, newDotAws); err != nil {
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
