package main

import (
	"fmt"
	"github.com/kingjh/baidu-ai-go-sdk/vision"
	"github.com/kingjh/baidu-ai-go-sdk/vision/image_classify"
)

const (
	// This Api Key and Api Secret is just for example,
	// you should get your own first.
	APIKEY    = "HXWyX72qTSYYngVWjo0LdOa4"
	APISECRET = "FmyLRIKLyqLan99HzfuuvRvK2gl6XGwR"
)

var client *image_classify.ImageClassifyClient

func init() {
	client = image_classify.NewImageClassifyClient(APIKEY, APISECRET)
}

func main() {
	Dish()
}

func Dish() {
	rs, err := client.Dish(
		vision.MustFromFile("D:\\projs\\go\\src\\github.com\\chenqinghe\\baidu-ai-go-sdk\\example\\vision\\image_classify\\JiuHuangChaoDan.jpg"),
	)
	if err != nil {
		panic(err)
	}

	fmt.Println(rs.ToString())
}
