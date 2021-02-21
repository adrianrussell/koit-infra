package vultr

import (
	"testing"
)

func TestWriteStartup(t *testing.T) {

	a := VultrImplementation{}

	got := a.WriteStartup()
	if got != "Provision Vultr!" {
		t.Errorf("WriteStartup = %s; want Provision Vultr!", got)
	}
}
