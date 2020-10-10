package engine

import "testing"

func TestNewEcho(t *testing.T) {
	e := NewEcho(EchoConfig{

	})

	e.Start(":9999")
}
