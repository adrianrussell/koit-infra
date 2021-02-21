package main

type status interface {
	GetInstances()
}

type topologyStatus struct {
	Status status
}
