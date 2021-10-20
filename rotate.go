package rotakey

import (
	"fmt"

	log "github.com/sirupsen/logrus"
)

// Execute rotates user's AWS credentials
// Possible error messages are:
// - LimitExceeded: Cannot exceed quota for AccessKeysPerUser: %d
// - failed to create access key due to %v
// - InvalidClientTokenId at CreateAccessKey: The security token (%s) included in the request is invalid
// - InvalidClientTokenId at DeleteAccessKey: The security token (%s) included in the request is invalid
// ...
// - failed to update access key (%s)
// - no profile with %s access key id
// - unable to find AWS profile due to SESSION_TOKEN
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
	newAwsCfg, err := resetConfig(newKeys)
	if err != nil {
		return err
	}
	updatedID, err := getAccessKeyID(newAwsCfg)
	if err != nil {
		return err
	} else if idAtStart == updatedID {
		return fmt.Errorf("failed to update access key (%s)", idAtStart)
	}
	log.Debug("new access key ID is " + updatedID)
	newIAM := newIAM(newAwsCfg)
	return newIAM.deleteAccessKey(idAtStart)
}
