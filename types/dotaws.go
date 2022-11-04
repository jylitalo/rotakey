package types

import (
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"gopkg.in/ini.v1"
)

type DotAws interface {
	Load() error
	GetProfile(accessKeyId string) (*ini.Section, error)
	Save(profile *ini.Section, accessKey types.AccessKey) error
}
