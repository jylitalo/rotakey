package rotakey

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	iamtypes "github.com/aws/aws-sdk-go-v2/service/iam/types"
	log "github.com/sirupsen/logrus"
	"golang.org/x/sys/unix"
	ini "gopkg.in/ini.v1"
)

type DotAws struct {
	filename string
	iniFile  *ini.File
}

type DotAwsOptions struct {
	filename string
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

func (da DotAws) Load() error {
	if da.filename == "" {
		fname, err := credentialsFile(config.DefaultSharedCredentialsFilename())
		if err != nil {
			return err
		}
		da.filename = fname
	}
	iniFile, err := ini.Load(da.filename)
	if err != nil {
		return err
	}
	da.iniFile = iniFile
	return nil
}

func (da DotAws) GetProfile(accessKeyId string) (*ini.Section, error) {
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

func (da DotAws) Save(profile *ini.Section, accessKey iamtypes.AccessKey) error {
	log.Info("REAL.da.Save")
	profile.Key("aws_access_key_id").SetValue(*accessKey.AccessKeyId)
	profile.Key("aws_secret_access_key").SetValue(*accessKey.SecretAccessKey)
	if err := da.iniFile.SaveTo(da.filename); err != nil {
		return fmt.Errorf("failed to save %s due to %v", da.filename, err.Error())
	}
	log.Debug("Updated config saved into " + da.filename)
	return nil
}
