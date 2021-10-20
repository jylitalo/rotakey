package rotakey

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"

	log "github.com/sirupsen/logrus"
	ini "gopkg.in/ini.v1"
)

type dotAws struct {
	filename string
	iniFile  *ini.File
}

func credentialsFile() (string, error) {
	fname := config.DefaultSharedCredentialsFilename()
	_, err := os.Stat(fname)
	switch {
	case err == nil:
		return fname, nil
	case os.IsNotExist(err):
		return fname, fmt.Errorf("%s does not exist", fname)
	default:
		return fname, fmt.Errorf("failed to stat %s", fname)
	}
}

func NewDotAws() (*dotAws, error) {
	fname, err := credentialsFile()
	if err != nil {
		return nil, err
	}
	iniFile, err := ini.Load(fname)
	if err != nil {
		return nil, err
	}
	return &dotAws{
		filename: fname,
		iniFile:  iniFile,
	}, nil
}

func (client *dotAws) getProfile(accessKeyId string) (*ini.Section, error) {
	for _, profile := range client.iniFile.Sections() {
		id, err := profile.GetKey("aws_access_key_id")
		if err != nil {
			continue
		}
		if id.String() == accessKeyId {
			log.Debugf("Found %s from %v profile", accessKeyId, profile.Name())
			return profile, nil
		}
	}
	if os.Getenv("SESSION_TOKEN") != "" {
		return nil, fmt.Errorf("unable to find AWS profile due to SESSION_TOKEN")
	}
	return nil, fmt.Errorf("no profile with %s access key id", accessKeyId)
}

func (client *dotAws) save(profile *ini.Section, accessKey *types.AccessKey) error {
	profile.Key("aws_access_key_id").SetValue(*accessKey.AccessKeyId)
	profile.Key("aws_secret_access_key").SetValue(*accessKey.SecretAccessKey)
	return client.iniFile.SaveTo(client.filename)
}
