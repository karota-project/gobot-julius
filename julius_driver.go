package julius

import (
  "github.com/hybridgroup/gobot"
)

type JuliusDriver struct {
  gobot.Driver
}

type JuliusInterface interface {
}

func NewJuliusDriver(a *JuliusAdaptor, name string) *JuliusDriver {
  return &JuliusDriver{
    Driver: *gobot.NewDriver(
      name,
      "julius.JuliusDriver",
      a,
    ),
  }
}

func (j *JuliusDriver) adaptor() *JuliusAdaptor {
  return j.Driver.Adaptor().(*JuliusAdaptor)
}

func (j *JuliusDriver) Start() bool { return true }
func (j *JuliusDriver) Halt() bool { return true }
