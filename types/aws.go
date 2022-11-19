package types

type AwsConfig interface {
	AccessKeyID() (string, error)
	LoadDefaultConfig() error
	NewIam() AwsIam
}
