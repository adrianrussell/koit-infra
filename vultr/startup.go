package vultr

type VultrImplementation struct {
}

func (implementation VultrImplementation) WriteStartup() string {
	return "Provision Vultr!"
}
