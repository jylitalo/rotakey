package rotakey

import "fmt"

type awsConfigMock struct {
	accessKey string
}

func newAwsConfigMock() (AwsConfigIface, error) {
	return awsConfigMock{accessKey: "AKIABCDEFGHIJKLKMNOP"}, nil
}

func (client awsConfigMock) accessKeyID() (string, error) {
	if client.accessKey == "" {
		return "", fmt.Errorf("no AWS access key")
	}
	return client.accessKey, nil
}

func (client awsConfigMock) newIam() awsIamIface {
	return awsIamMock{callback: &client}
}
