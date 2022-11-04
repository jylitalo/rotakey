package rotakey

import (
	"context"
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
)

type awsIam struct {
	sdk *iam.Client
}

func (ia *awsIam) CreateAccessKey() (*types.AccessKey, error) {
	resp, err := ia.sdk.CreateAccessKey(context.TODO(), &iam.CreateAccessKeyInput{})
	if err == nil {
		return resp.AccessKey, nil
	}
	txt := err.Error()
	switch {
	case strings.Contains(txt, "LimitExceeded: Cannot exceed quota"):
		start := strings.Index(txt, "LimitExceeded: Cannot exceed quota")
		return nil, fmt.Errorf(txt[start:])
	case strings.Contains(txt, "InvalidClientTokenId: The security token included in the request is invalid"):
		return nil, fmt.Errorf("InvalidClientTokenId at CreateAccessKey: The security token included in the request is invalid")
	}
	return nil, fmt.Errorf("failed to create access key due to %s", txt)
}

func (ia *awsIam) DeleteAccessKey(accessKeyId string) error {
	_, err := ia.sdk.DeleteAccessKey(context.TODO(), &iam.DeleteAccessKeyInput{AccessKeyId: &accessKeyId})
	switch {
	case err == nil:
		return err
	case strings.Contains(err.Error(), "InvalidClientTokenId: The security token included in the request is invalid"):
		return fmt.Errorf("InvalidClientTokenId at DeleteAccessKey: The security token (%s) included in the request is invalid", accessKeyId)
	default:
		return fmt.Errorf("failed to delete access key due to %v", err)
	}
}
