package mock

import (
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go-v2/service/iam/types"
)

type AwsIam struct {
	AwsAccessKeyId      string
	FailCreateAccessKey int
}

func (ia *AwsIam) CreateAccessKey() (*types.AccessKey, error) {
	accessKeyId := ia.AwsAccessKeyId
	if len(accessKeyId) < 2 {
		return nil, fmt.Errorf("AccessKey (%s) is too short", accessKeyId)
	}
	if strings.HasSuffix(accessKeyId, "CreateERR") {
		return nil, fmt.Errorf("error condition triggered")
	}
	if ia.FailCreateAccessKey > 0 {
		ia.FailCreateAccessKey--
		return nil, fmt.Errorf("InvalidClientTokenId at CreateAccessKey: The security token included in the request is invalid")
	}
	ia.AwsAccessKeyId = accessKeyId[:len(accessKeyId)-2] + "Z"
	secretAccessKey := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMN"
	awsConfigAccessKey = accessKeyId
	return &types.AccessKey{
		AccessKeyId:     &accessKeyId,
		SecretAccessKey: &secretAccessKey,
	}, nil
}

func (ia *AwsIam) DeleteAccessKey(accessKeyId string) error {
	return nil
}
