package image_classify

import "github.com/kingjh/baidu-ai-go-sdk"

type ImageClassifyClient struct {
	*gosdk.Client
}

func NewImageClassifyClient(apiKey, secretKey string) *ImageClassifyClient {
	return &ImageClassifyClient{
		Client: gosdk.NewClient(apiKey, secretKey),
	}
}
