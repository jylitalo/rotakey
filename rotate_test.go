package rotakey

import "testing"

func TestExecute(t *testing.T) {
	err := NewExec().Execute(ExecuteInput{NewAwsConfig: newAwsConfigMock, NewDotAws: newDotAwsMock})
	if err != nil {
		t.Errorf("Execute failed due to %v", err)
	}
}
