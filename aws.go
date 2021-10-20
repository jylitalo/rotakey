package rotakey

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

func getConfig() (aws.Config, error) {
	return config.LoadDefaultConfig(context.TODO())
}

func getAccessKeyID(cfg aws.Config) (string, error) {
	creds, err := cfg.Credentials.Retrieve(context.TODO())
	return creds.AccessKeyID, err
}
