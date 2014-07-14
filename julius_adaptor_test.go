package julius

import (
	"github.com/hybridgroup/gobot"
	"testing"
)

func initTestJuliusAdaptor() *JuliusAdaptor {
	return NewJuliusAdaptor("myAdaptor")
}

func TestJuliusAdaptorConnect(t *testing.T) {
	a := initTestJuliusAdaptor()
	gobot.Expect(t, a.Connect(), true)
}

func TestJuliusAdaptorFinalize(t *testing.T) {
	a := initTestJuliusAdaptor()
	gobot.Expect(t, a.Finalize(), true)
}
