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

	print(a)
	print(b)
}

func print(s provision.Startup) {
	fmt.Println(s.WriteStartup())
}

func configureInstance(s provision.InfrastructureDefinition) {
	s.Define().AddInstance().Number(1).Size("Small")
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
