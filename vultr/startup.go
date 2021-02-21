package vultr

import "fmt"

type VultrImplementation struct{}

func (implementation *VultrImplementation) WriteStartup() {
	fmt.Println("Provision Vultr!")
}
