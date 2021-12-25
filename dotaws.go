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

type DotAwsIface interface {
	getProfile(accessKeyId string) (*ini.Section, error)
	save(profile *ini.Section, accessKey *types.AccessKey) error
}

func credentialsFile(fname string) (string, error) {
	_, err := os.Stat(fname)
	switch {
	case err == nil:
	case os.IsNotExist(err):
		err = fmt.Errorf("%s does not exist", fname)
	default:
		err = fmt.Errorf("failed to check %s due to %v", fname, err)
	}
	return fname, err
}

func newDotAws() (DotAwsIface, error) {
	if fname, err := credentialsFile(config.DefaultSharedCredentialsFilename()); err != nil {
		return nil, err
	} else if iniFile, err := ini.Load(fname); err != nil {
		return nil, err
	} else {
		return dotAws{filename: fname, iniFile: iniFile}, nil
	}
}

func (client dotAws) getProfile(accessKeyId string) (*ini.Section, error) {
	for _, profile := range client.iniFile.Sections() {
		id, err := profile.GetKey("aws_access_key_id")
		if err == nil && id.String() == accessKeyId {
			log.Infof("Found %s from %v profile", accessKeyId, profile.Name())
			return profile, nil
		}
	}
	if os.Getenv("AWS_SESSION_TOKEN") != "" {
		return nil, fmt.Errorf("unable to find AWS profile due to AWS_SESSION_TOKEN")
	}
	return nil, fmt.Errorf("no profile with %s access key id", accessKeyId)
}

func (client dotAws) save(profile *ini.Section, accessKey *types.AccessKey) error {
	profile.Key("aws_access_key_id").SetValue(*accessKey.AccessKeyId)
	profile.Key("aws_secret_access_key").SetValue(*accessKey.SecretAccessKey)
	if err := client.iniFile.SaveTo(client.filename); err != nil {
		return fmt.Errorf("failed to save %s due to %v", client.filename, err.Error())
	}
	return nil
}
