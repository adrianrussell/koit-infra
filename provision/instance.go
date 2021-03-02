package provision

type Instance interface {
	Size(string) Instance
	Number(int) Instance
}

type InfrastructureTypes interface {
	AddInstance() Instance
}

type InfrastructureDefinition interface {
	Define() InfrastructureTypes
}
