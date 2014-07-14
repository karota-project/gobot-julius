package julius

import (
	"github.com/hybridgroup/gobot"
	"net"
)

type JuliusAdaptor struct {
	gobot.Adaptor
	tcpConn *net.TCPConn
}

func NewJuliusAdaptor(name string, port string) *JuliusAdaptor {
	return &JuliusAdaptor{
		Adaptor: *gobot.NewAdaptor(
			name,
			"julius.JuliusAdaptor",
			port,
		),
	}
}

func (j *JuliusAdaptor) Connect() bool {
	if j.Connected() == true {
		disconnect(j)
	}
	connect(j)
	return true
}

func (j *JuliusAdaptor) Finalize() bool {
	disconnect(j)
	return true
}

func connect(j *JuliusAdaptor) {
	addr, err := net.ResolveTCPAddr("tcp", j.Port())
	if err != nil {
		panic(err)
	}

	conn, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		panic(err)
	}

	j.tcpConn = conn
	j.SetConnected(true)
}

func disconnect(j *JuliusAdaptor) {
	j.tcpConn.Close()
	j.SetConnected(false)
}
