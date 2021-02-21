package aws

import (
	"testing"
)

func TestWriteStartup(t *testing.T) {

	a := AWSImplementation{}

	got := a.WriteStartup()
	if got != "Provision AWS!" {
		t.Errorf("WriteStartup = %s; want Provision AWS!", got)
	}
}
