package rotakey

var awsConfigMockAccessKey = "AKIABCDEFGHIJKLKMNOP"

type awsConfigMock struct {
	failCreateAccessKey int
}

func newAwsConfigMock() (AwsConfigIface, error) {
	return &awsConfigMock{}, nil
}

func (cf *awsConfigMock) accessKeyID() (string, error) {
	return awsConfigMockAccessKey, nil
}

func (cf *awsConfigMock) newIam() awsIam {
	return &awsIamMock{
		failCreateAccessKey: &cf.failCreateAccessKey,
	}
}
