package rotakey

import (
	"os"
	"testing"

	log "github.com/sirupsen/logrus"
	ini "gopkg.in/ini.v1"

	"github.com/jylitalo/rotakey/mock"
	"github.com/jylitalo/rotakey/types"
)

// Scenarios:
// - all is well (TestExecute)
// - no access key defined (TestAwsConfigMissing)
// - credentials file not found (TestDotAwsMissing)
// - create access key pair fails for unspecified reason (TestCreateAccessKeyError)
// - create access key failed because user already has two key pairs

func TestExecute(t *testing.T) {
	rot := &Rotate{}
	err := rot.Execute(&mock.AwsConfig{AwsAccessKeyId: "AKIABCDEFGHIJKLKMNOP"}, &mock.DotAws{})
	if err != nil {
		t.Errorf("Execute failed due to %v", err)
	}
}

func TestExecuteWithOneFailure(t *testing.T) {
	rot := &Rotate{}
	err := rot.Execute(
		&mock.AwsConfig{AwsAccessKeyId: "AKIABCDEFGHIJKLKMNOP", FailCreateAccessKey: 1}, &mock.DotAws{})
	if err != nil {
		t.Errorf("ExecuteWithOneFailure failed due to %v", err)
	}
}

func TestAwsConfigMissing(t *testing.T) {
	rot := &Rotate{}
	err := rot.Execute(&AwsConfig{}, &mock.DotAws{})
	if err == nil {
		t.Errorf("Execute did't abort due to err")
	}
	log.Info(err)
}

func newDotAwsMissing() (types.DotAws, error) {
	fname, _ := os.CreateTemp(".", "invalid-*")
	os.Remove(fname.Name())
	if fname, err := credentialsFile(fname.Name()); err != nil {
		return nil, err
	} else if iniFile, err := ini.Load(fname); err != nil {
		return nil, err
	} else {
		return DotAws{filename: fname, iniFile: iniFile}, nil
	}
}
func TestDotAwsMissing(t *testing.T) {
	rot := &Rotate{}
	err := rot.Execute(&mock.AwsConfig{AwsAccessKeyId: "AKIABCDEFGHIJKLKMNOP"}, &mock.DotAws{})
	if err == nil {
		t.Errorf("Execute did't abort due to err")
	}
}

func TestCreateAccesskeyError(t *testing.T) {
	rot := &Rotate{}
	err := rot.Execute(&mock.AwsConfig{AwsAccessKeyId: "AKIABCDEFGHCreateERR"}, &mock.DotAws{})
	if err == nil {
		t.Errorf("Execute did't abort due to err")
	}
}
