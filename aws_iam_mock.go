package rotakey

import (
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go-v2/service/iam/types"
)

type awsIamMock struct {
	callback *awsConfigMock
}

func (client awsIamMock) createAccessKey() (*types.AccessKey, error) {
	accessKeyId, _ := client.callback.accessKeyID()
	if strings.HasSuffix(accessKeyId, "Z") {
		return nil, fmt.Errorf("error condition triggered")
	}
	accessKeyId = accessKeyId[:len(accessKeyId)-2] + "Z"
	secretAccessKey := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMN"
	return &types.AccessKey{
		AccessKeyId:     &accessKeyId,
		SecretAccessKey: &secretAccessKey,
	}, nil
}

func (client awsIamMock) deleteAccessKey(accessKeyId string) error {
	return nil
}
