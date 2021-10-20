package rotakey

import (
	"context"
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
)

type awsIAM struct {
	sdk *iam.Client
	cfg aws.Config
}

func newIAM(cfg aws.Config) *awsIAM {
	return &awsIAM{sdk: iam.NewFromConfig(cfg), cfg: cfg}
}

func (client *awsIAM) createAccessKey() (*types.AccessKey, error) {
	resp, err := client.sdk.CreateAccessKey(context.TODO(), nil)
	switch {
	case err == nil:
		return resp.AccessKey, nil
	case strings.Contains(err.Error(), "LimitExceeded: Cannot exceed quota"):
		start := strings.Index(err.Error(), "LimitExceeded: Cannot exceed quota")
		return nil, fmt.Errorf(err.Error()[start:])
	case strings.Contains(err.Error(), "InvalidClientTokenId: The security token included in the request is invalid"):
		return nil, fmt.Errorf("InvalidClientTokenId at CreateAccessKey: The security token included in the request is invalid")
	default:
		return nil, fmt.Errorf("failed to create access key due to %v", err)
	}
}

func (client *awsIAM) deleteAccessKey(accessKeyId string) error {
	_, err := client.sdk.DeleteAccessKey(context.TODO(), &iam.DeleteAccessKeyInput{AccessKeyId: &accessKeyId})
	switch {
	case err == nil:
		return err
	case strings.Contains(err.Error(), "InvalidClientTokenId: The security token included in the request is invalid"):
		return fmt.Errorf("InvalidClientTokenId at DeleteAccessKey: The security token (%s) included in the request is invalid", accessKeyId)
	default:
		return fmt.Errorf("failed to delete access key due to %v", err)
	}
}
