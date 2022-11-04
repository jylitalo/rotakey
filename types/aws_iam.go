package types

import (
	iamtypes "github.com/aws/aws-sdk-go-v2/service/iam/types"
)

type AwsIam interface {
	CreateAccessKey() (*iamtypes.AccessKey, error)
	DeleteAccessKey(accessKeyId string) error
}
