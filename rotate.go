package rotakey

import (
	"fmt"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	log "github.com/sirupsen/logrus"
)

type ExecuteInput struct {
	NewAwsConfig func() (AwsConfigIface, error)
	NewDotAws    func() (DotAwsIface, error)
}

type Exec struct{}

type ExecIface interface {
	Execute(ExecuteInput) error
}

func NewExec() ExecIface {
	return Exec{}
}

func robustCreateAccessKey(awsCfg AwsConfigIface) (*types.AccessKey, error) {
	var err error
	for attempt := 1; attempt < 5; attempt++ {
		var newKeys *types.AccessKey
		iam := awsCfg.newIam()
		newKeys, err = iam.createAccessKey()
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
func (client Exec) Execute(params ExecuteInput) error {
	// setup
	if params.NewAwsConfig == nil {
		params.NewAwsConfig = newAwsConfig
	}
	if params.NewDotAws == nil {
		params.NewDotAws = newDotAws
	}
	awsCfg, errA := params.NewAwsConfig()
	fileCfg, errB := params.NewDotAws()
	if err := CoalesceError(errA, errB); err != nil {
		return fmt.Errorf("constructors in execute failed due to %s", err.Error())
	}
	idAtStart, errA := awsCfg.accessKeyID()
	profile, errB := fileCfg.getProfile(idAtStart)
	if profile != nil {
		log.Debugf("Found access key (%s) from %s profile", idAtStart, profile.Name())
	}
	if err := CoalesceError(errA, errB); err != nil {
		return fmt.Errorf("checking access key and profile failed due to %s", err.Error())
	}
	newKeys, err := robustCreateAccessKey(awsCfg)
	if err != nil {
		return err
	}
	errA = fileCfg.save(profile, *newKeys) // verify
	newAwsCfg, errB := params.NewAwsConfig()
	if err := CoalesceError(errA, errB); err != nil {
		return fmt.Errorf("changes in execute failed due to %s", err.Error())
	}
	updatedID, err := newAwsCfg.accessKeyID()
	if err != nil {
		return fmt.Errorf("failed to get new access key (%s) due to %s", *newKeys.AccessKeyId, err.Error())
	} else if idAtStart == updatedID {
		return fmt.Errorf("failed to update access key (%s)", idAtStart)
	}
	log.Infof("New access key created (id: %s)", *newKeys.AccessKeyId)
	log.Infof("Deleted old access key (id: %s)", idAtStart)
	return awsCfg.newIam().deleteAccessKey(idAtStart)
}

func CoalesceError(errs ...error) error {
	for _, err := range errs {
		if err != nil {
			return err
		}
	}
	return nil
}
