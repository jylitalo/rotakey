package rotakey

import "testing"

func TestExecute(t *testing.T) {
	err := NewExec().Execute(ExecuteInput{NewAwsConfig: newAwsConfigMock, NewDotAws: newDotAwsMock})
	if err != nil {
		t.Errorf("Execute failed due to %v", err)
	}
}

func newAwsConfigMockWithErr() (AwsConfigIface, error) {
	return awsConfigMock{accessKey: "AKIABCDEFGHIJKLKMNOZ"}, nil
}

func TestCreateAccesskeyError(t *testing.T) {
	err := NewExec().Execute(ExecuteInput{NewAwsConfig: newAwsConfigMockWithErr, NewDotAws: newDotAwsMock})
	if err == nil {
		t.Errorf("Execute did't abort due to err")
	}
}