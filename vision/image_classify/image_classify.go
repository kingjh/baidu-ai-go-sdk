package image_classify

import (
	"errors"
	"github.com/kingjh/baidu-ai-go-sdk/vision"
)

const (
	IMAGE_CLASSIFY_ADVANCED_GENERAL_URL = "https://aip.baidubce.com/rest/2.0/image-classify/v2/advanced_general"
	IMAGE_CLASSIFY_DISH_URL             = "https://aip.baidubce.com/rest/2.0/image-classify/v2/dish"
	IMAGE_CLASSIFY_LOGO_URL             = "https://aip.baidubce.com/rest/2.0/image-classify/v2/logo"
	IMAGE_CLASSIFY_ANIMAL_URL           = "https://aip.baidubce.com/rest/2.0/image-classify/v1/animal"
	IMAGE_CLASSIFY_PLANT_URL            = "https://aip.baidubce.com/rest/2.0/image-classify/v1/plant"
	IMAGE_CLASSIFY_FLOWER_URL           = "https://aip.baidubce.com/rest/2.0/image-classify/v1/flower"
	IMAGE_CLASSIFY_INGREDIENT_URL       = "https://aip.baidubce.com/rest/2.0/image-classify/v1/classify/ingredient"
	IMAGE_CLASSIFY_LANDMARK_URL         = "https://aip.baidubce.com/rest/2.0/image-classify/v1/landmark"
	IMAGE_CLASSIFY_CAR_URL              = "https://aip.baidubce.com/rest/2.0/image-classify/v1/car"
)

//细粒度图像识别——人物/场景识别
func (oc *ImageClassifyClient) AdvancedGeneral(image interface{}, params ...RequestParam) (*ImageClassifyResponse, error) {
	return oc.imageClassify(image, IMAGE_CLASSIFY_ADVANCED_GENERAL_URL, defaultDishParams, params...)
}

//细粒度图像识别——菜式识别
func (oc *ImageClassifyClient) Dish(image interface{}, params ...RequestParam) (*ImageClassifyResponse, error) {
	return oc.imageClassify(image, IMAGE_CLASSIFY_DISH_URL, defaultDishParams, params...)
}

//细粒度图像识别——商标识别
func (oc *ImageClassifyClient) Logo(image interface{}, params ...RequestParam) (*ImageClassifyResponse, error) {
	return oc.imageClassify(image, IMAGE_CLASSIFY_LOGO_URL, defaultDishParams, params...)
}

//细粒度图像识别——动物识别
func (oc *ImageClassifyClient) Animal(image interface{}, params ...RequestParam) (*ImageClassifyResponse, error) {
	return oc.imageClassify(image, IMAGE_CLASSIFY_ANIMAL_URL, defaultDishParams, params...)
}

//细粒度图像识别——植物识别
func (oc *ImageClassifyClient) Plant(image interface{}, params ...RequestParam) (*ImageClassifyResponse, error) {
	return oc.imageClassify(image, IMAGE_CLASSIFY_PLANT_URL, defaultDishParams, params...)
}

//细粒度图像识别——花卉识别
func (oc *ImageClassifyClient) Flower(image interface{}, params ...RequestParam) (*ImageClassifyResponse, error) {
	return oc.imageClassify(image, IMAGE_CLASSIFY_FLOWER_URL, defaultDishParams, params...)
}

//细粒度图像识别——果蔬识别
func (oc *ImageClassifyClient) Ingredient(image interface{}, params ...RequestParam) (*ImageClassifyResponse, error) {
	return oc.imageClassify(image, IMAGE_CLASSIFY_INGREDIENT_URL, defaultDishParams, params...)
}

//细粒度图像识别——地标识别
func (oc *ImageClassifyClient) Landmark(image interface{}, params ...RequestParam) (*ImageClassifyResponse, error) {
	return oc.imageClassify(image, IMAGE_CLASSIFY_LANDMARK_URL, defaultDishParams, params...)
}

//细粒度图像识别——车型识别
func (oc *ImageClassifyClient) Car(image interface{}, params ...RequestParam) (*ImageClassifyResponse, error) {
	return oc.imageClassify(image, IMAGE_CLASSIFY_CAR_URL, defaultDishParams, params...)
}

func (oc *ImageClassifyClient) imageClassify(image interface{}, url string, def map[string]interface{}, params ...RequestParam) (*ImageClassifyResponse, error) {
	requestParams, err := parseRequestParam(image, def, params...)
	if err != nil {
		return nil, err
	}

	return oc.doRequest(url, requestParams)
}

func parseRequestParam(image interface{}, def map[string]interface{}, params ...RequestParam) (map[string]interface{}, error) {
	image, ok := image.(string)
	if !ok {
		// image is *vision.Image
		image, _ := image.(*vision.Image)
		if image.Reader == nil {
			if image.Url == "" {
				return nil, errors.New("image source is empty")
			} else {
				def["url"] = image.Url
				delete(def, "image")
			}
		} else {
			base64Str, err := image.Base64Encode()
			if err != nil {
				return nil, err
			}
			def["image"] = base64Str
		}
	} else {
		def["image"] = image
	}

	for _, fn := range params {
		fn(def)
	}

	return def, nil

}
