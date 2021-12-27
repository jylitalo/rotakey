package rotakey

var awsConfigMockAccessKey = "AKIABCDEFGHIJKLKMNOP"

type awsConfigMock struct {
	failCreateAccessKey int
}

func newAwsConfigMock() (AwsConfigIface, error) {
	return &awsConfigMock{}, nil
}

func (client *awsConfigMock) accessKeyID() (string, error) {
	return awsConfigMockAccessKey, nil
}

func (client *awsConfigMock) newIam() awsIamIface {
	return &awsIamMock{
		failCreateAccessKey: &client.failCreateAccessKey,
	}
}
