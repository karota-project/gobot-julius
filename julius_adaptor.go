package julius

import (
  "github.com/hybridgroup/gobot"
)

type JuliusAdaptor struct {
  gobot.Adaptor
}

func NewJuliusAdaptor(name string) *JuliusAdaptor {
  return &JuliusAdaptor{
    Adaptor: *gobot.NewAdaptor(
      name,
      "julius.JuliusAdaptor",
    ),
  }
}

func (j *JuliusAdaptor) Connect() bool {
  return true
}

func (j *JuliusAdaptor) Finalize() bool {
  return true
}
