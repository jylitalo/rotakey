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
func Execute() error {
	// setup
	awsCfg, err := getConfig()
	if err != nil {
		return err
	}
	idAtStart, err := getAccessKeyID(awsCfg)
	if err != nil {
		return err
	}
	fileCfg, err := NewDotAws()
	if err != nil {
		return err
	}
	// validate
	profile, err := fileCfg.getProfile(idAtStart)
	if err != nil {
		return err
	}
	// execute
	iam := newIAM(awsCfg)
	newKeys, err := iam.createAccessKey()
	if err != nil {
		return err
	} else if err = fileCfg.save(profile, newKeys); err != nil {
		return err
	}
	// verify
	newAwsCfg, err := getConfig()
	if err != nil {
		return err
	}
	updatedID, err := getAccessKeyID(newAwsCfg)
	if err != nil {
		return err
	} else if idAtStart == updatedID {
		return fmt.Errorf("failed to update access key (%s)", idAtStart)
	}
	log.Debug("new AWS access key ID is " + updatedID)
	return iam.deleteAccessKey(idAtStart)
}
