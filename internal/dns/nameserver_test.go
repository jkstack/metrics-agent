package dns

import "testing"

func TestNameServer(t *testing.T) {
	nameServers, err := NameServer()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(nameServers)
}
