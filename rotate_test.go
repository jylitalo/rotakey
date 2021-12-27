package rotakey

import (
	"io/ioutil"
	"os"
	"testing"

	ini "gopkg.in/ini.v1"
)

// Scenarios:
// - all is well (TestExecute)
// - no access key defined (TestAwsConfigMissing)
// - credentials file not found (TestDotAwsMissing)
// - create access key pair fails for unspecified reason (TestCreateAccessKeyError)
// - create access key failed because user already has two key pairs

func TestExecute(t *testing.T) {
	awsConfigMockAccessKey = "AKIABCDEFGHIJKLKMNOP"
	err := NewExec().Execute(ExecuteInput{NewAwsConfig: newAwsConfigMock, NewDotAws: newDotAwsMock})
	if err != nil {
		t.Errorf("Execute failed due to %v", err)
	}
}

func TestAwsConfigMissing(t *testing.T) {
	awsConfigMockAccessKey = ""
	err := NewExec().Execute(ExecuteInput{NewAwsConfig: newAwsConfigMock, NewDotAws: newDotAws})
	if err == nil {
		t.Errorf("Execute did't abort due to err")
	}
}

func newDotAwsMissing() (DotAwsIface, error) {
	fname, _ := ioutil.TempFile(".", "invalid-*")
	os.Remove(fname.Name())
	if fname, err := credentialsFile(fname.Name()); err != nil {
		return nil, err
	} else if iniFile, err := ini.Load(fname); err != nil {
		return nil, err
	} else {
		return dotAws{filename: fname, iniFile: iniFile}, nil
	}
}
func TestDotAwsMissing(t *testing.T) {
	awsConfigMockAccessKey = "AKIABCDEFGHIJKLKMNOP"
	err := NewExec().Execute(ExecuteInput{NewAwsConfig: newAwsConfigMock, NewDotAws: newDotAwsMissing})
	if err == nil {
		t.Errorf("Execute did't abort due to err")
	}
}

func TestCreateAccesskeyError(t *testing.T) {
	awsConfigMockAccessKey = "AKIABCDEFGHCreateERR"
	err := NewExec().Execute(ExecuteInput{NewAwsConfig: newAwsConfigMock, NewDotAws: newDotAws})
	if err == nil {
		t.Errorf("Execute did't abort due to err")
	}
}
