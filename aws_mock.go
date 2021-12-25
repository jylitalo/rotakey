package rotakey

type awsConfigMock struct{}

func newAwsConfigMock() (AwsConfigIface, error) {
	return awsConfigMock{}, nil
}

func (client awsConfigMock) accessKeyID() (string, error) {
	return "AKIABCDEFGHIJKLKMNOP", nil
}

func (client awsConfigMock) newIam() awsIamIface {
	return awsIamMock{}
}
