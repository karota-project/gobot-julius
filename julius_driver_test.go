package julius

import (
  "github.com/hybridgroup/gobot"
  "testing"
)

func initTestJuliusDriver() *JuliusDriver {
  return NewJuliusDriver(NewJuliusAdaptor("myAdaptor"), "myDriver")
}

func TestJuliusDriverStart(t *testing.T) {
  d := initTestJuliusDriver()
  gobot.Expect(t, d.Start(), true)
}

func TestJuliusDriverHalt(t *testing.T) {
  d := initTestJuliusDriver()
  gobot.Expect(t, d.Halt(), true)
}
