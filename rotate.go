package rotakey

import (
	"fmt"
	"strings"
	"time"

	iamtypes "github.com/aws/aws-sdk-go-v2/service/iam/types"
	log "github.com/sirupsen/logrus"

	"github.com/jylitalo/rotakey/internal"
	"github.com/jylitalo/rotakey/types"
)

type Rotate struct{}

func robustCreateAccessKey(awsCfg types.AwsConfig) (*iamtypes.AccessKey, error) {
	var err error
	iam := awsCfg.NewIam()
	for attempt := 1; attempt < 5; attempt++ {
		var newKeys *iamtypes.AccessKey
		newKeys, err = iam.CreateAccessKey()
		switch {
		case err == nil:
			return newKeys, nil
		case strings.HasPrefix(err.Error(), "InvalidClientTokenId at CreateAccessKey"):
			log.Debugf("Failed attempt #%d due to InvalidClientTokenId", attempt)
			time.Sleep(time.Second * time.Duration(attempt))
		default:
			return nil, fmt.Errorf("access key creation failed due to %s", err.Error())
		}
	}
	return nil, err
}

// Execute rotates user's AWS credentials in ~/.aws/credentials
// Possible error messages are:
// - %s does not exist
// - failed to check %s
// - failed to create access key due to %v
// - failed to delete access key due to %v
// - failed to save %s due to %v
// - no profile with %s access key id
// - unable to find AWS profile due to AWS_SESSION_TOKEN
// - InvalidClientTokenId at CreateAccessKey: The security token (%s) included in the request is invalid
// - InvalidClientTokenId at DeleteAccessKey: The security token (%s) included in the request is invalid
// - LimitExceeded: Cannot exceed quota for AccessKeysPerUser: %d
func (client Rotate) Execute(awsCfg types.AwsConfig, fileCfg types.DotAws) error {
	// setup
	_ = awsCfg.LoadDefaultConfig()
	idAtStart, errA := awsCfg.AccessKeyID()
	_ = fileCfg.Load()
	profile, errB := fileCfg.GetProfile(idAtStart)
	if err := internal.CoalesceError(errA, errB); err != nil {
		return fmt.Errorf("checking access key and profile failed due to %s", err.Error())
	}
	if profile != nil {
		log.Debugf("Found access key (%s) from %s profile", idAtStart, profile.Name())
	}
	newKeys, err := robustCreateAccessKey(awsCfg)
	if err != nil {
		return err
	}
	log.Info("Going for save")
	if err = fileCfg.Save(profile, *newKeys); err != nil {
		return err
	}
	log.Info("Going for verify")
	// verify
	_ = awsCfg.LoadDefaultConfig()
	updatedID, err := awsCfg.AccessKeyID()
	if err != nil {
		return fmt.Errorf("failed to get new access key (%s) due to %s", *newKeys.AccessKeyId, err.Error())
	} else if idAtStart == updatedID {
		return fmt.Errorf("failed to update access key (%s)", idAtStart)
	}
	log.Infof("New access key created (id: %s)", *newKeys.AccessKeyId)
	log.Infof("Deleted old access key (id: %s)", idAtStart)
	return awsCfg.NewIam().DeleteAccessKey(idAtStart)
}
