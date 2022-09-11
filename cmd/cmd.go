package cmd

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/jylitalo/rotakey"
)

type NewCmdInput struct {
	Use          string
	NewExec      func() rotakey.ExecIface
	NewAwsConfig func() (rotakey.AwsConfig, error)
	NewDotAws    func() (rotakey.DotAws, error)
}

func NewCmd(params NewCmdInput) *cobra.Command {
	if params.Use == "" {
		params.Use = "rotakey"
	}
	cmd := &cobra.Command{
		Use:   params.Use,
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
			if params.NewExec == nil {
				params.NewExec = rotakey.NewExec
			}
			err := params.NewExec().Execute(rotakey.ExecuteInput{
				NewAwsConfig: params.NewAwsConfig,
				NewDotAws:    params.NewDotAws,
			})
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
