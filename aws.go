package rotakey

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/iam"
)

type AwsConfig interface {
	accessKeyID() (string, error)
	newIam() awsIam
}

type awsConfigImpl struct {
	config aws.Config
}

func newAwsConfig() (AwsConfig, error) {
	var err error
	cfg := &awsConfigImpl{}
	cfg.config, err = config.LoadDefaultConfig(context.TODO())
	return cfg, err
}

func (cf *awsConfigImpl) accessKeyID() (string, error) {
	creds, err := cf.config.Credentials.Retrieve(context.TODO())
	return creds.AccessKeyID, err
}

func (cf *awsConfigImpl) newIam() awsIam {
	return &awsIamImpl{sdk: iam.NewFromConfig(cf.config)}
}
