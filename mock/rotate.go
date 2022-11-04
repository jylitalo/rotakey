package mock

import "github.com/jylitalo/rotakey/types"

type Rotate struct{}

func (rk Rotate) Execute(awsCfg types.AwsConfig, fileCfg types.DotAws) error {
	return nil
}
