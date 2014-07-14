package julius

import (
	"encoding/xml"
	"github.com/hybridgroup/gobot"
	"regexp"
	"strings"
)

const (
	START_PROC    = "StartProc"
	END_PROC      = "EndProc"
	START_RECOG   = "StartRecog"
	END_RECOG     = "EndRecog"
	INPUT         = "Input"
	INPUT_PARAM   = "InputParam"
	GMM           = "Gmm"
	RECOG_OUT     = "RecogOut"
	RECOG_FAIL    = "RecogFail"
	REJECTED      = "Rejected"
	GRAPH_OUT     = "GraphOut"
	GRAM_INFO     = "GramInfo"
	SYS_INFO      = "SysInfo"
	ENGINE_INFO   = "EngineInfo"
	GRAMMER       = "Grammer"
	RECOG_PROCESS = "RecogProcess"
)

type JuliusDriver struct {
	gobot.Driver
}

type JuliusInterface interface {
}

func NewJuliusDriver(a *JuliusAdaptor, name string) *JuliusDriver {
	s := &JuliusDriver{
		Driver: *gobot.NewDriver(
			name,
			"julius.JuliusDriver",
			a,
		),
	}

	s.AddEvent(START_PROC)
	s.AddEvent(END_PROC)
	s.AddEvent(START_RECOG)
	s.AddEvent(END_RECOG)
	s.AddEvent(INPUT)
	s.AddEvent(INPUT_PARAM)
	s.AddEvent(GMM)
	s.AddEvent(RECOG_OUT)
	s.AddEvent(RECOG_FAIL)
	s.AddEvent(REJECTED)
	s.AddEvent(GRAPH_OUT)
	s.AddEvent(GRAM_INFO)
	s.AddEvent(SYS_INFO)
	s.AddEvent(ENGINE_INFO)
	s.AddEvent(GRAMMER)
	s.AddEvent(RECOG_PROCESS)

	/* add julius control command
	s.AddCommand("", func(params map[string]interface{}) interface{} {

	})
	*/

	return s
}

func (j *JuliusDriver) adaptor() *JuliusAdaptor {
	return j.Driver.Adaptor().(*JuliusAdaptor)
}

func (j *JuliusDriver) Start() bool {
	go func() {
		for {
			msg := make([]byte, 10240) //FIXME バッファーオーバーフローする可能性あり
			rlen, err := j.adaptor().tcpConn.Read(msg)

			if err != nil {
				panic(err)
			}

			// 不正な値(<s>タグ)が入ってきてGo言語のXMLパーサでパースできないので、空白で置換しているため値は入らない
			re, _ := regexp.Compile("CLASSID=\".*?\" ")
			s := re.ReplaceAllString(string(msg[:rlen]), "")
			sa := strings.Split(s, ".\n")

			for _, ss := range sa {
				// StartProc
				startProc := StartProc{}
				err = xml.Unmarshal([]byte(ss), &startProc)
				if err == nil {
					gobot.Publish(j.Event(START_PROC), startProc)
					continue
				}

				// EndProc
				endProc := EndProc{}
				err = xml.Unmarshal([]byte(ss), &endProc)
				if err == nil {
					gobot.Publish(j.Event(END_PROC), endProc)
					continue
				}

				// StartRecog
				startRecog := StartRecog{}
				err = xml.Unmarshal([]byte(ss), &startRecog)
				if err == nil {
					gobot.Publish(j.Event(START_RECOG), startRecog)
					continue
				}

				// EndRecog
				endRecog := EndRecog{}
				err = xml.Unmarshal([]byte(ss), &endRecog)
				if err == nil {
					gobot.Publish(j.Event(END_RECOG), endRecog)
					continue
				}

				// Input
				input := Input{}
				err = xml.Unmarshal([]byte(ss), &input)
				if err == nil {
					gobot.Publish(j.Event(INPUT), input)
					continue
				}

				// InputParam
				inputParam := InputParam{}
				err = xml.Unmarshal([]byte(ss), &inputParam)
				if err == nil {
					gobot.Publish(j.Event(INPUT_PARAM), inputParam)
					continue
				}

				// Gmm
				gmm := Gmm{}
				err = xml.Unmarshal([]byte(ss), &gmm)
				if err == nil {
					gobot.Publish(j.Event(GMM), gmm)
					continue
				}

				// RecogOut
				recogOut := RecogOut{}
				err = xml.Unmarshal([]byte(ss), &recogOut)
				if err == nil {
					gobot.Publish(j.Event(RECOG_OUT), recogOut)
					continue
				}

				// RecogFail
				recogFail := RecogFail{}
				err = xml.Unmarshal([]byte(ss), &recogFail)
				if err == nil {
					gobot.Publish(j.Event(RECOG_FAIL), recogFail)
					continue
				}

				// Rejected
				rejected := Rejected{}
				err = xml.Unmarshal([]byte(ss), &rejected)
				if err == nil {
					gobot.Publish(j.Event(REJECTED), rejected)
					continue
				}

				// GraphOut
				graphOut := GraphOut{}
				err = xml.Unmarshal([]byte(ss), &graphOut)
				if err == nil {
					gobot.Publish(j.Event(GRAPH_OUT), graphOut)
					continue
				}

				// GramInfo
				gramInfo := GramInfo{}
				err = xml.Unmarshal([]byte(ss), &gramInfo)
				if err == nil {
					gobot.Publish(j.Event(GRAM_INFO), gramInfo)
					continue
				}

				// SysInfo
				sysInfo := SysInfo{}
				err = xml.Unmarshal([]byte(ss), &sysInfo)
				if err == nil {
					gobot.Publish(j.Event(SYS_INFO), sysInfo)
					continue
				}

				// EngineInfo
				engineInfo := EngineInfo{}
				err = xml.Unmarshal([]byte(ss), &engineInfo)
				if err == nil {
					gobot.Publish(j.Event(ENGINE_INFO), engineInfo)
					continue
				}

				// Grammer
				grammer := Grammer{}
				err = xml.Unmarshal([]byte(ss), &grammer)
				if err == nil {
					gobot.Publish(j.Event(GRAMMER), engineInfo)
					continue
				}

				// RecogProcess
				recogProcess := RecogProcess{}
				err = xml.Unmarshal([]byte(ss), &recogProcess)
				if err == nil {
					gobot.Publish(j.Event(RECOG_PROCESS), engineInfo)
					continue
				}
			}
		}
	}()

	return true
}

func (j *JuliusDriver) Halt() bool {
	return true
}
