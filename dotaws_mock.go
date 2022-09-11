package rotakey

import (
	"github.com/aws/aws-sdk-go-v2/service/iam/types"

	ini "gopkg.in/ini.v1"
)

type dotAwsMock struct{}

func newDotAwsMock() (DotAwsIface, error) {
	return dotAwsMock{}, nil
}

func (da dotAwsMock) getProfile(accessKeyId string) (*ini.Section, error) {
	accessKeyID := "AKIABCDEFGHIJKLKMNOZ"
	secretAccessKey := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMZ"
	file := ini.Empty()
	section, errA := file.NewSection("mock")
	_, errB := section.NewKey("aws_access_key_id", accessKeyID)
	_, errC := section.NewKey("aws_secret_access_key", secretAccessKey)
	return section, CoalesceError(errA, errB, errC)
}

func (da dotAwsMock) save(profile *ini.Section, accessKey types.AccessKey) error {
	return nil
}
