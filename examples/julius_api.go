package main

import (
	"fmt"

	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/api"
	"github.com/karota-project/gobot-julius"
)

func main() {
	master := gobot.NewGobot()
	api.NewAPI(master).Start()

	juliusAdaptor := julius.NewJuliusAdaptor("julius-a01", "localhost:10500")
	juliusDriver := julius.NewJuliusDriver(juliusAdaptor, "julius-d01")

	master.AddRobot(
		gobot.NewRobot(
			"julius",
			[]gobot.Connection{juliusAdaptor},
			[]gobot.Device{juliusDriver},
			func() {
				fmt.Println("work")

				events := []string{
					julius.START_PROC,
					julius.END_PROC,
					julius.START_RECOG,
					julius.END_RECOG,
					julius.INPUT,
					julius.INPUT_PARAM,
					julius.GMM,
					julius.RECOG_OUT,
					julius.RECOG_FAIL,
					julius.REJECTED,
					julius.GRAPH_OUT,
					julius.GRAM_INFO,
					julius.SYS_INFO,
					julius.ENGINE_INFO,
					julius.GRAMMER,
					julius.RECOG_PROCESS,
				}

				for _, event := range events {
					gobot.On(juliusDriver.Event(event), func(data interface{}) {
						fmt.Println("-----")
						fmt.Println(data)
						fmt.Println("-----")
					})
				}
			}))

	master.Start()
}
