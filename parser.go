package julius

import (
	"encoding/xml"
)

//認識エンジン動作開始
//<STARTPROC/>
type StartProc struct {
	XMLName xml.Name `xml:"STARTPROC"`
}

//認識エンジン停止時
//<ENDPROC/>
type EndProc struct {
	XMLName xml.Name `xml:"ENDPROC"`
}

//認識処理開始
//<STARTRECOG/>
type StartRecog struct {
	XMLName xml.Name `xml:"STARTRECOG"`
}

//認識処理終了
//<ENDRECOG/>
type EndRecog struct {
	XMLName xml.Name `xml:"ENDRECOG"`
}

//入力開始
//<INPUT STATUS="LISTEN" TIME="..."/>
//入力始端検知
//<INPUT STATUS="STARTREC" TIME="..."/>
//入力終端検知
//<INPUT STATUS="ENDREC" TIME="..."/>
type Input struct {
	XMLName xml.Name `xml:"INPUT"`
	Status  string   `xml:"STATUS,attr"`
	Time    string   `xml:"TIME,attr"`
}

//入力長情報
//<INPUTPARAM FRAMES="..." MSEC="..."/>
type InputParam struct {
	XMLName xml.Name `xml:"INPUT"`
	Frames  string   `xml:"FRAMES,attr"`
	Msec    string   `xml:"MSEC,attr"`
}

//GMM最尤のモデルとそのGMM信頼度
//<GMM RESULT="xxx" CMSCORE="..."/>
type Gmm struct {
	XMLName xml.Name `xml:"GMM"`
}

//認識結果（成功時）
//<RECOGOUT>...</RECOGOUT>
type RecogOut struct {
	XMLName xml.Name `xml:"RECOGOUT"`
	Shypo   Shypo    `xml:"SHYPO"`
}

type Shypo struct {
	XMLName xml.Name `xml:"SHYPO"`
	Rank    string   `xml:"RANK,attr"`
	Score   string   `xml:"SCORE,attr"`
	Whypo   []Whypo  `xml:"WHYPO"`
}

type Whypo struct {
	XMLName xml.Name `xml:"WHYPO"`
	Word    string   `xml:"WORD,attr"`
	ClassId string   `xml:"CLASSID,attr"` // 不正な値(<s>タグ)が入ってきてGo言語のXMLパーサでパースできないので、空白で置換しているため値は入らない
	Phone   string   `xml:"PHONE,attr"`
	Cm      string   `xml:"CM,attr"`
}

//認識失敗
//<RECOGFAIL/>
type RecogFail struct {
	XMLName xml.Name `xml:"RECOGFAIL"`
}

//入力棄却
//<REJECTED REASON="...">
type Rejected struct {
	XMLName xml.Name `xml:"REJECTED"`
	Reason  string   `xml:"REASON,attr"`
}

//単語グラフ
//<GRAPHOUT>...</GRAPHOUT>
type GraphOut struct {
	XMLName xml.Name `xml:"GRAPHOUT"`
}

//現在エンジンが保持している文法情報
//<GRAMINFO>...</GRAMINFO>
type GramInfo struct {
	XMLName xml.Name `xml:"GRAMINFO"`
}

//エンジンの現在の状態
//<SYSINFO PROCESS="ACTIVE|SLEEP">
type SysInfo struct {
	XMLName xml.Name `xml:"SYSINFO"`
	Process string   `xml:"PROCESS"`
}

//エンジンのバージョン
//<ENGINEINFO TYPE="Julius" VERSION="4.1" CONF="fast"/>
type EngineInfo struct {
	XMLName xml.Name `xml:"ENGINEINFO"`
	Type    string   `xml:"TYPE,attr"`
	Version string   `xml:"VERSION,attr"`
	Config  string   `xml:"CONF,attr"`
}

//文法受け取り確認
//<GRAMMAR STATUS="RECEIVED"/>
//文法準備完了確認
//<GRAMMAR STATUS="READY"/>
//エラー
//<GRAMMAR STATUS="ERROR" REASON="..."/>
type Grammer struct {
	XMLName xml.Name `xml:"GRAMMAR"`
	Status  string   `xml:"STATUS"`
	Reason  string   `xml:"REASON"`
}

//認識処理インスタンスの情報
//<RECOGPROCESS>...<RECOGPROCESS/>
type RecogProcess struct {
	XMLName xml.Name `xml:"RECOGPROCESS"`
}
