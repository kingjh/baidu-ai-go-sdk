package voice

import (
	"encoding/base64"

	"net"

	"github.com/imroc/req"
)

const ASR_URL = "http://vop.baidu.com/server_api"

const FILE_TYPE_PCM = "pcm"
const FILE_TYPE_WAV = "wav"
const FILE_TYPE_AMR = "amr"
const FILE_TYPE_MP3 = "mp3"

const DEV_PID_MANDARIN_MIX = 1536
const DEV_PID_MANDARIN_PURE = 1537
const DEV_PID_ENGLISH = 1737
const DEV_PID_CANTONESE = 1637
const DEV_PID_SICHUAN_DIALECT = 1837
const DEV_PID_MANDARIN_FAR_FIELD = 1936

//语音识别响应信息
type ASRResponse struct {
	*req.Resp
}

//语音识别参数
type ASRParams struct {
	Format  string `json:"format"`  //语音的格式，pcm 或者 wav 或者 amr。不区分大小写
	Rate    int    `json:"rate"`    //采样率，支持 8000 或者 16000
	Channel int    `json:"channel"` //声道数，仅支持单声道，请填写固定值 1
	Cuid    string `json:"cuid"`    //用户唯一标识，用来区分用户，计算UV值。建议填写能区分用户的机器 MAC 地址或 IMEI 码，长度为60字符以内
	Token   string `json:"token"`   //开放平台获取到的开发者access_token
	DevPid  int    `json:"dev_pid"` //识别语言: 1536=普通话(支持简单的英文识别)|1537=普通话(纯中文识别)|1737=英语|1637=粤语|1837=四川话|1936=普通话远场
	Speech  string `json:"speech"`  //真实的语音数据 ，需要进行base64 编码。与len参数连一起使用
	Length  int    `json:"len"`     //原始语音长度，单位字节
}

type ASRParam func(params *ASRParams)

func Format(fmt string) ASRParam {

	if fmt != "pcm" && fmt != "wav" && fmt != "amr" {
		fmt = "pcm"
	}
	return func(params *ASRParams) {
		params.Format = fmt
	}
}

func Rate(rate int) ASRParam {
	if rate != 8000 && rate != 16000 {
		rate = 16000
	}
	return func(params *ASRParams) {
		params.Rate = rate
	}
}

func Channel(c int) ASRParam {
	return func(params *ASRParams) {
		params.Channel = 1 //固定值1
	}
}

func DevPid(devPid int) ASRParam {
	if devPid != DEV_PID_MANDARIN_MIX && devPid != DEV_PID_MANDARIN_PURE && devPid != DEV_PID_ENGLISH && devPid != DEV_PID_CANTONESE && devPid != DEV_PID_SICHUAN_DIALECT && devPid != DEV_PID_MANDARIN_FAR_FIELD {
		devPid = DEV_PID_MANDARIN_MIX
	}
	return func(params *ASRParams) {
		params.DevPid = devPid
	}
}

// SpeechToText 语音识别，将语音翻译成文字
func (vc *VoiceClient) SpeechToText(voiceBytes []byte, params ...ASRParam) (*ASRResponse, error) {
	if err := vc.Auth(); err != nil {
		return nil, err
	}

	spch := base64.StdEncoding.EncodeToString(voiceBytes)

	var cuid string
	netitfs, err := net.Interfaces()
	if err != nil {
		cuid = "anonymous"
	} else {
		for _, itf := range netitfs {
			if cuid = itf.HardwareAddr.String(); len(cuid) > 0 {
				break
			}
		}
	}

	asrParams := &ASRParams{
		Format:  "pcm",
		Rate:    16000,
		Channel: 1,
		Cuid:    cuid,
		Token:   vc.AccessToken,
		DevPid:  DEV_PID_MANDARIN_MIX,
		Speech:  spch,
		Length:  len(voiceBytes),
	}

	for _, fn := range params {
		fn(asrParams)
	}

	header := req.Header{
		"Content-Type": "application/json",
	}

	resp, err := req.Post(ASR_URL, header, req.BodyJSON(asrParams))
	if err != nil {
		return nil, err
	}

	return &ASRResponse{
		Resp: resp,
	}, nil

}
