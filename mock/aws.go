package mock

import "github.com/jylitalo/rotakey/types"

var awsConfigAccessKey = "AKIABCDEFGHIJKLKMNOP"

type AwsConfig struct {
	AwsAccessKeyId      string
	FailCreateAccessKey int
}

func (cf *AwsConfig) AccessKeyID() (string, error) {
	return awsConfigAccessKey, nil
}

func (cf *AwsConfig) NewIam() types.AwsIam {
	return &AwsIam{
		AwsAccessKeyId:      cf.AwsAccessKeyId,
		FailCreateAccessKey: cf.FailCreateAccessKey,
	}
}
func (cf *AwsConfig) LoadDefaultConfig() error {
	return nil
}
