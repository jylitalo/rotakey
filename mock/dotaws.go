package mock

import (
	"fmt"

	iamtypes "github.com/aws/aws-sdk-go-v2/service/iam/types"
	log "github.com/sirupsen/logrus"
	ini "gopkg.in/ini.v1"

	"github.com/jylitalo/rotakey/internal"
)

type DotAws struct{}

const DefaultAccessKey = "AKIABCDEFGHIJKLKMNOP"

func (da DotAws) Load() error { return nil }
func (da DotAws) GetProfile(searchKeyID string) (*ini.Section, error) {
	daAccessKeyID := DefaultAccessKey
	if daAccessKeyID != searchKeyID {
		return nil, fmt.Errorf("no profile with %s access key id", searchKeyID)
	}
	secretAccessKey := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMZ"
	file := ini.Empty()
	section, errA := file.NewSection("mock")
	_, errB := section.NewKey("aws_access_key_id", daAccessKeyID)
	_, errC := section.NewKey("aws_secret_access_key", secretAccessKey)
	return section, internal.CoalesceError(errA, errB, errC)
}

func (da DotAws) Save(profile *ini.Section, accessKey iamtypes.AccessKey) error {
	log.Info("mock.da.Save")
	return nil
}
