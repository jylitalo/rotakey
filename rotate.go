package rotakey

import (
	"fmt"

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
		return err
	}
	idAtStart, errA := awsCfg.accessKeyID()
	iam := awsCfg.newIam()
	// validate
	profile, errB := fileCfg.getProfile(idAtStart)
	if profile != nil {
		log.Debugf("Found access key (%s) from %s profile", idAtStart, profile.Name())
	}
	// execute
	newKeys, errC := iam.createAccessKey()
	if err := CoalesceError(errA, errB, errC); err != nil {
		return err
	}
	errA = fileCfg.save(profile, *newKeys) // verify
	newAwsCfg, errB := newAwsConfig()
	if err := CoalesceError(errA, errB); err != nil {
		return err
	}
	updatedID, err := newAwsCfg.accessKeyID()
	if err != nil {
		return err
	} else if idAtStart == updatedID {
		return fmt.Errorf("failed to update access key (%s)", idAtStart)
	}
	log.Infof("New access key created (id: %s)", *newKeys.AccessKeyId)
	log.Infof("Deleted old access key (id: %s)", idAtStart)
	return iam.deleteAccessKey(idAtStart)
}

func CoalesceError(errs ...error) error {
	for _, err := range errs {
		if err != nil {
			return err
		}
	}
	return nil
}
