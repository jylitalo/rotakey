package rotakey

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/iam"

	"github.com/jylitalo/rotakey/types"
)

type AwsConfig struct {
	config aws.Config
}

func (cf *AwsConfig) AccessKeyID() (string, error) {
	creds, err := cf.config.Credentials.Retrieve(context.TODO())
	return creds.AccessKeyID, err
}

func (cf *AwsConfig) LoadDefaultConfig() error {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return err
	}
	cf.config = cfg
	return nil
}

func (cf *AwsConfig) NewIam() types.AwsIam {
	return &awsIam{sdk: iam.NewFromConfig(cf.config)}
}
