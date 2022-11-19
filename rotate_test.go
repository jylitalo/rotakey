package rotakey

import (
	"os"
	"testing"

	log "github.com/sirupsen/logrus"

	"github.com/jylitalo/rotakey/mock"
)

// Scenarios:
// - all is well (TestExecute)
// - no access key defined (TestAwsConfigMissing)
// - credentials file not found (TestDotAwsMissing)
// - create access key pair fails for unspecified reason (TestCreateAccessKeyError)
// - create access key failed because user already has two key pairs

func TestExecute(t *testing.T) {
	rot := &Rotate{}
	err := rot.Execute(&mock.AwsConfig{AwsAccessKeyId: mock.DefaultAccessKey}, &mock.DotAws{})
	if err != nil {
		t.Errorf("Execute failed due to %v", err)
	}
}

func TestExecuteWithOneFailure(t *testing.T) {
	rot := &Rotate{}
	err := rot.Execute(
		&mock.AwsConfig{AwsAccessKeyId: mock.DefaultAccessKey, FailCreateAccessKey: 1}, &mock.DotAws{})
	if err != nil {
		t.Errorf("ExecuteWithOneFailure failed due to %v", err)
	}
}

func TestAwsConfigMissing(t *testing.T) {

	rot := &Rotate{}
	err := rot.Execute(&mock.AwsConfig{}, &mock.DotAws{})
	if err == nil {
		t.Errorf("Execute did't abort due to err")
	}
	log.Info(err)
}

func TestDotAwsMissing(t *testing.T) {
	tmpFile, _ := os.CreateTemp(".", "invalid-*")
	os.Remove(tmpFile.Name())
	dot := DotAws{filename: tmpFile.Name()}

	rot := &Rotate{}
	err := rot.Execute(&mock.AwsConfig{AwsAccessKeyId: mock.DefaultAccessKey}, dot)
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
