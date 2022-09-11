package rotakey

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"golang.org/x/sys/unix"

	log "github.com/sirupsen/logrus"
	ini "gopkg.in/ini.v1"
)

type dotAwsImpl struct {
	filename string
	iniFile  *ini.File
}

type DotAws interface {
	getProfile(accessKeyId string) (*ini.Section, error)
	save(profile *ini.Section, accessKey types.AccessKey) error
}

func credentialsFile(fname string) (string, error) {
	_, err := os.Stat(fname)
	switch {
	case err == nil:
		switch {
		case unix.Access(fname, unix.R_OK) != nil:
			err = fmt.Errorf("no read access to %s due to %v", fname, unix.Access(fname, unix.R_OK))
		case unix.Access(fname, unix.W_OK) != nil:
			err = fmt.Errorf("no write access to %s due to %v", fname, unix.Access(fname, unix.W_OK))
		}
	case os.IsNotExist(err):
		err = fmt.Errorf("%s does not exist", fname)
	default:
		err = fmt.Errorf("failed to check %s due to %v", fname, err)
	}
	return fname, err
}

func newDotAws() (DotAws, error) {
	if fname, err := credentialsFile(config.DefaultSharedCredentialsFilename()); err != nil {
		return nil, err
	} else if iniFile, err := ini.Load(fname); err != nil {
		return nil, err
	} else {
		return dotAwsImpl{filename: fname, iniFile: iniFile}, nil
	}
}

func (da dotAwsImpl) getProfile(accessKeyId string) (*ini.Section, error) {
	for _, profile := range da.iniFile.Sections() {
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

func (da dotAwsImpl) save(profile *ini.Section, accessKey types.AccessKey) error {
	profile.Key("aws_access_key_id").SetValue(*accessKey.AccessKeyId)
	profile.Key("aws_secret_access_key").SetValue(*accessKey.SecretAccessKey)
	if err := da.iniFile.SaveTo(da.filename); err != nil {
		return fmt.Errorf("failed to save %s due to %v", da.filename, err.Error())
	}
	log.Debug("Updated config saved into " + da.filename)
	return nil
}
