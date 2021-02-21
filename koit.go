package main

import (
	"fmt"
	"koit/aws"
	"koit/provision"
	"koit/vultr"
)

func main() {
	fmt.Println("Starting Provisioning")

	a := aws.AWSImplementation{}

	b := vultr.VultrImplementation{}

	print(&a)
	print(&b)
}

func print(s provision.Startup) {
	s.WriteStartup()
}

type topology struct {
	Name                  string
	ComputeResources      []compute
	NetworkResources      []network
	LoadBalancerResources []loadBalancer
}

type compute struct {
	Name string
}

type network struct {
	Name string
}

type loadBalancer struct {
	Name string
}
