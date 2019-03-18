package image_classify

import "github.com/imroc/req"

type ImageClassifyResponse struct {
	*req.Resp
}

func (oc *ImageClassifyClient) doRequest(url string, params map[string]interface{}) (response *ImageClassifyResponse, err error) {

	if err := oc.Auth(); err != nil {
		return nil, err
	}

	header := req.Header{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	url += "?access_token=" + oc.AccessToken

	resp, err := req.Post(url, req.Param(params), header)
	if err != nil {
		return nil, err
	}
	return &ImageClassifyResponse{
		Resp: resp,
	}, nil
}
