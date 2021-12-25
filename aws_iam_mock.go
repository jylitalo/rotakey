package rotakey

import (
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
)

type awsIamMock struct{}

func (client awsIamMock) createAccessKey() (*types.AccessKey, error) {
	accessKeyId := "AKIABCDEFGHIJKLKMNOP"
	secretAccessKey := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMN"
	return &types.AccessKey{
		AccessKeyId:     &accessKeyId,
		SecretAccessKey: &secretAccessKey,
	}, nil
}

func (client awsIamMock) deleteAccessKey(accessKeyId string) error {
	return nil
}
