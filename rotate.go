package rotakey

import (
	"fmt"

	log "github.com/sirupsen/logrus"
)

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
func Execute(newAwsConfig func() (AwsConfigIface, error), newDotAws func() (DotAwsIface, error)) error {
	// setup
	awsCfg, errA := newAwsConfig()
	fileCfg, errB := newDotAws()
	if err := CoalesceError(errA, errB); err != nil {
		return err
	}
	idAtStart, errA := awsCfg.accessKeyID()
	iam := awsCfg.newIam()
	// validate
	profile, errB := fileCfg.getProfile(idAtStart)
	log.Debugf("Found access key (%s) from %s profile", idAtStart, profile)
	// execute
	newKeys, errC := iam.createAccessKey()
	if err := CoalesceError(errA, errB, errC); err != nil {
		return err
	}
	errA = fileCfg.save(profile, newKeys) // verify
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
