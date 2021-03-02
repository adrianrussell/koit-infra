package aws

type AWSImplementation struct {
}

func (implementation AWSImplementation) WriteStartup() string {
	return "Provision AWS!"
}
