package rotakey

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/iam"
)

type AwsConfigIface interface {
	accessKeyID() (string, error)
	newIam() awsIam
}

type awsConfig struct {
	config aws.Config
}

func newAwsConfig() (AwsConfigIface, error) {
	var err error
	cfg := &awsConfig{}
	cfg.config, err = config.LoadDefaultConfig(context.TODO())
	return cfg, err
}

func (cf *awsConfig) accessKeyID() (string, error) {
	creds, err := cf.config.Credentials.Retrieve(context.TODO())
	return creds.AccessKeyID, err
}

func (cf *awsConfig) newIam() awsIam {
	return &awsIamImpl{sdk: iam.NewFromConfig(cf.config)}
}
