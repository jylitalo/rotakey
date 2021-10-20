package rotakey

import (
	"context"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/aws/aws-sdk-go-v2/service/sts"
)

func getConfig() (aws.Config, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	return cfg, err
}

func resetConfig(accessKey *types.AccessKey) (aws.Config, error) {
	return config.LoadDefaultConfig(context.TODO(),
		// Hard-coded credentials.
		config.WithCredentialsProvider(credentials.StaticCredentialsProvider{
			Value: aws.Credentials{
				AccessKeyID: *accessKey.AccessKeyId, SecretAccessKey: *accessKey.SecretAccessKey,
				SessionToken: "reset",
				Source:       "hard-coded credentials",
			},
		}),
	)
}

func getAccessKeyID(cfg aws.Config) (string, error) {
	creds, err := cfg.Credentials.Retrieve(context.TODO())
	return creds.AccessKeyID, err
}

func getUsername(cfg aws.Config) (string, error) {
	resp, err := sts.NewFromConfig(cfg).GetCallerIdentity(context.TODO(), &sts.GetCallerIdentityInput{})
	if err != nil {
		return "", err
	}
	fields := strings.Split(*resp.Arn, "/")
	return fields[1], nil
}
