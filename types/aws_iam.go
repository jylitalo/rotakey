package types

import (
	"context"

	iamtypes "github.com/aws/aws-sdk-go-v2/service/iam/types"
)

type AwsIam interface {
	CreateAccessKey(ctx context.Context) (*iamtypes.AccessKey, error)
	DeleteAccessKey(ctx context.Context, accessKeyId string) error
}
