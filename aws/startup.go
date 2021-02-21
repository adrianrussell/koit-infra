package aws

import "fmt"

type AWSImplementation struct{}

func (implementation *AWSImplementation) WriteStartup() {
	fmt.Println("Provision AWS!")
}
