package types

type Rotate interface {
	Execute(awsCfg AwsConfig, fileCfg DotAws) error
}
