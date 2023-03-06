package rotakey

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/iam"

	awstypes "github.com/jylitalo/rotakey/types"
)

type AwsConfig struct {
	config aws.Config
}

func (cf *AwsConfig) AccessKeyID(ctx context.Context) (string, error) {
	creds, err := cf.config.Credentials.Retrieve(ctx)
	return creds.AccessKeyID, err
}

func (cf *AwsConfig) LoadDefaultConfig(ctx context.Context) error {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return err
	}
	cf.config = cfg
	return nil
}

func (cf *AwsConfig) NewIam() awstypes.AwsIam {
	return &awsIam{sdk: iam.NewFromConfig(cf.config)}
}
