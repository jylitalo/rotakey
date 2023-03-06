package types

import "context"

type AwsConfig interface {
	AccessKeyID(ctx context.Context) (string, error)
	LoadDefaultConfig(ctx context.Context) error
	NewIam() AwsIam
}
