package rotakey

import (
	"github.com/aws/aws-sdk-go-v2/service/iam/types"

	ini "gopkg.in/ini.v1"
)

type dotAwsMock struct{}

func newDotAwsMock() (DotAwsIface, error) {
	return dotAwsMock{}, nil
}

func (client dotAwsMock) getProfile(accessKeyId string) (*ini.Section, error) {
	accessKeyID := "AKIABCDEFGHIJKLKMNOZ"
	secretAccessKey := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMZ"
	file := ini.Empty()
	section, _ := file.NewSection("mock")
	section.NewKey("aws_access_key_id", accessKeyID)
	section.NewKey("aws_secret_access_key", secretAccessKey)
	return section, nil
}

func (client dotAwsMock) save(profile *ini.Section, accessKey *types.AccessKey) error {
	return nil
}
