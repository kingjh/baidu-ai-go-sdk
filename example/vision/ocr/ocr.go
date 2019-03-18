package main

import (
	"fmt"
	"github.com/kingjh/baidu-ai-go-sdk/vision"
	"github.com/kingjh/baidu-ai-go-sdk/vision/ocr"
)

const (
	// This Api Key and Api Secret is just for example,
	// you should get your own first.
	APIKEY    = "OyqxWZRpQhMLfPdPnG68DBFc"
	APISECRET = "TmcSKBQ70itvxSpkjH42r3e52jnp1oQx"
)

var client *ocr.OCRClient

func init() {
	client = ocr.NewOCRClient(APIKEY, APISECRET)
}

func main() {
	GeneralBasicRecognize()
	//GeneralRecognizeEnhanced()
}

func GeneralBasicRecognize() {
	rs, err := client.GeneralBasicRecognize(
		vision.MustFromFile("D:\\projs\\go\\src\\github.com\\chenqinghe\\baidu-ai-go-sdk\\example\\vision\\ocr\\ocr.jpg"),
		ocr.DetectDirection(),       //是否检测图像朝向，默认不检测
		ocr.DetectLanguage(),        //是否检测语言，默认不检测。
		ocr.LanguageType("CHN_ENG"), //识别语言类型，默认为CHN_ENG。
		ocr.WithProbability(),       //是否返回识别结果中每一行的置信度
	)
	if err != nil {
		panic(err)
	}

	fmt.Println(rs.ToString())
}

//func GeneralRecognizeEnhanced() {
//
//	resp, err := client.GeneralRecognizeEnhanced(
//		vision.MustFromFile("ocr.jpg"),
//		ocr.DetectDirection(),
//		ocr.DetectLanguage(),
//		ocr.LanguageType("CHN_ENG"),
//	)
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println(resp.ToString())
//
//}
