package mock

import (
	"context"
	"fmt"
	"strings"

	iamtypes "github.com/aws/aws-sdk-go-v2/service/iam/types"

	"github.com/jylitalo/rotakey/types"
)

type AwsConfig struct {
	AwsAccessKeyId      string
	FailCreateAccessKey int
}

func (cf *AwsConfig) AccessKeyID(ctx context.Context) (string, error) { return cf.AwsAccessKeyId, nil }
func (cf *AwsConfig) NewIam() types.AwsIam                            { return cf }
func (cf *AwsConfig) LoadDefaultConfig(ctx context.Context) error     { return nil }

// IAM part

func (cf *AwsConfig) CreateAccessKey(ctx context.Context) (*iamtypes.AccessKey, error) {
	accessKeyId := cf.AwsAccessKeyId
	if len(accessKeyId) < 2 {
		return nil, fmt.Errorf("AccessKey (%s) is too short", accessKeyId)
	}
	if strings.HasSuffix(accessKeyId, "CreateERR") {
		return nil, fmt.Errorf("error condition triggered")
	}
	if cf.FailCreateAccessKey > 0 {
		cf.FailCreateAccessKey--
		return nil, fmt.Errorf("InvalidClientTokenId at CreateAccessKey: The security token included in the request is invalid")
	}
	cf.AwsAccessKeyId = accessKeyId[:len(accessKeyId)-2] + "Z"
	secretAccessKey := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMN"
	return &iamtypes.AccessKey{
		AccessKeyId:     &cf.AwsAccessKeyId,
		SecretAccessKey: &secretAccessKey,
	}, nil
}

func (cf *AwsConfig) DeleteAccessKey(ctx context.Context, accessKeyId string) error { return nil }
