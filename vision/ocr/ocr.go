package ocr

import (
	"encoding/json"
	"errors"
	"github.com/kingjh/baidu-ai-go-sdk/vision"
)

const (
	OCR_GENERAL_BASIC_URL    = "https://aip.baidubce.com/rest/2.0/ocr/v1/general_basic"
	OCR_HANDWRITING_URL      = "https://aip.baidubce.com/rest/2.0/ocr/v1/handwriting"
	OCR_IDCARD_URL           = "https://aip.baidubce.com/rest/2.0/ocr/v1/idcard"
	OCR_BANKCARD_URL         = "https://aip.baidubce.com/rest/2.0/ocr/v1/bankcard"
	OCR_BUSINESS_LICENSE_URL = "https://aip.baidubce.com/rest/2.0/ocr/v1/business_license"
	OCR_PASSPORT_URL         = "https://aip.baidubce.com/rest/2.0/ocr/v1/passport"
	OCR_FORM_URL             = "https://aip.baidubce.com/rest/2.0/solution/v1/form_ocr/request"
	OCR_FORM_RESULT_URL      = "https://aip.baidubce.com/rest/2.0/solution/v1/form_ocr/get_request_result"
	OCR_RECEIPT_URL          = "https://aip.baidubce.com/rest/2.0/ocr/v1/receipt"
	OCR_VAT_INVOICE_URL      = "https://aip.baidubce.com/rest/2.0/ocr/v1/vat_invoice"
	OCR_TRAIN_TICKET_URL     = "https://aip.baidubce.com/rest/2.0/ocr/v1/train_ticket"
	OCR_TAX_RECEIPT_URL      = "https://aip.baidubce.com/rest/2.0/ocr/v1/taxi_receipt"
	OCR_DRIVING_LICENSE_URL  = "https://aip.baidubce.com/rest/2.0/ocr/v1/driving_license"
	OCR_VEHICLE_LICENSE_URL  = "https://aip.baidubce.com/rest/2.0/ocr/v1/vehicle_license"
	OCR_LICENSE_PLATE_URL    = "https://aip.baidubce.com/rest/2.0/ocr/v1/license_plate"
	OCR_VIN_CODE_URL         = "https://aip.baidubce.com/rest/2.0/ocr/v1/vin_code"
	OCR_NUMBERS_URL          = "https://aip.baidubce.com/rest/2.0/ocr/v1/numbers"
	OCR_WEBIMAGE_URL         = "https://aip.baidubce.com/rest/2.0/ocr/v1/webimage"

	//OCR_GENERAL_WITH_LOCATION_URL = "https://aip.baidubce.com/rest/2.0/ocr/v1/general"
	//OCR_GENERAL_ENHANCED_URL      = "https://aip.baidubce.com/rest/2.0/ocr/v1/general_enhanced"

	OCR_TYPE_LANGUAGE_CHN_ENG = "CHN_ENG"
	OCR_TYPE_LANGUAGE_ENG     = "ENG"
	OCR_TYPE_LANGUAGE_POR     = "POR"
	OCR_TYPE_LANGUAGE_FRE     = "FRE"
	OCR_TYPE_LANGUAGE_GER     = "GER"
	OCR_TYPE_LANGUAGE_ITA     = "ITA"
	OCR_TYPE_LANGUAGE_SPA     = "SPA"
	OCR_TYPE_LANGUAGE_RUS     = "RUS"
	OCR_TYPE_LANGUAGE_JAP     = "JAP"
	OCR_TYPE_LANGUAGE_KOR     = "KOR"

	OCR_TYPE_RESULT_EXCEL = "excel"
	OCR_TYPE_RESULT_JSON  = "json"
)

//GeneralBasicRecognize 通用文字识别
//识别图片中的文字信息
func (oc *OCRClient) GeneralBasicRecognize(image interface{}, params ...RequestParam) (*OCRResponse, error) {
	return oc.ocr(image, OCR_GENERAL_BASIC_URL, defaultGeneralBasicParams, params...)
}

//HandwritingRecognize 手写文字识别
func (oc *OCRClient) HandwritingRecognize(image interface{}, params ...RequestParam) (*OCRResponse, error) {
	return oc.ocr(image, OCR_HANDWRITING_URL, defaultHandwritingParams, params...)
}

//IdcardRecognize 身份证识别
//识别身份证正反面的文字信息
func (oc *OCRClient) IdcardRecognize(image interface{}, params ...RequestParam) (*OCRResponse, error) {
	return oc.ocr(image, OCR_IDCARD_URL, defaultIdcardParams, params...)
}

//BankcardRecognize 银行卡识别
//识别银行卡的卡号并返回发卡行和卡片性质信息
func (oc *OCRClient) BankcardRecognize(image interface{}, params ...RequestParam) (*OCRResponse, error) {
	return oc.ocr(image, OCR_BANKCARD_URL, defaultBankcardParams, params...)
}

//BusinessLicenseRecognize 营业执照识别
func (oc *OCRClient) BusinessLicenseRecognize(image interface{}, params ...RequestParam) (*OCRResponse, error) {
	return oc.ocr(image, OCR_BUSINESS_LICENSE_URL, defaultBusinessLicenseParams, params...)
}

//PassportRecognize 护照识别
func (oc *OCRClient) PassportRecognize(image interface{}, params ...RequestParam) (*OCRResponse, error) {
	return oc.ocr(image, OCR_PASSPORT_URL, defaultPassportParams, params...)
}

//FormRecognize 表格文字识别
//自动识别表格线及表格内容，结构化输出表头、表尾及每个单元格的文字内容
func (oc *OCRClient) FormRecognize(image interface{}, params ...RequestParam) (*OCRResponse, error) {
	resp, err := oc.ocr(image, OCR_FORM_URL, defaultFormParams, params...)
	if err != nil {
		return resp, err
	}

	jsonStr := resp.String()
	in := []byte(jsonStr)
	var raw map[string]interface{}
	json.Unmarshal(in, &raw)

	tmpInterfaces := raw["result"].([]interface{})
	tmpMap := tmpInterfaces[0].(map[string]interface{})

	percent := 0.0
	for {
		resp, err = oc.ocr(image, OCR_FORM_RESULT_URL, defaultFormResultParams, func(m map[string]interface{}) {
			m["request_id"] = tmpMap["request_id"]
		})
		if err != nil {
			return resp, err
		}

		jsonStr = resp.String()
		in = []byte(jsonStr)
		var rawLoop map[string]interface{}
		json.Unmarshal(in, &rawLoop)

		if rawLoop["result"] != nil {
			tmpMap = rawLoop["result"].(map[string]interface{})
			percent = tmpMap["percent"].(float64)
			if percent >= 100 {
				// 当进度>=100%时，才返回
				return resp, nil
			}
		}
	}

}

//ReceiptRecognize 通用票据识别
func (oc *OCRClient) ReceiptRecognize(image interface{}, params ...RequestParam) (*OCRResponse, error) {
	return oc.ocr(image, OCR_RECEIPT_URL, defaultReceiptParams, params...)
}

//VatInvoiceRecognize 增值税发票识别
func (oc *OCRClient) VatInvoiceRecognize(image interface{}, params ...RequestParam) (*OCRResponse, error) {
	return oc.ocr(image, OCR_VAT_INVOICE_URL, defaultVatInvoiceParams, params...)
}

//TrainTicketRecognize 火车票识别
func (oc *OCRClient) TrainTicketRecognize(image interface{}, params ...RequestParam) (*OCRResponse, error) {
	return oc.ocr(image, OCR_TRAIN_TICKET_URL, defaultTrainTicketParams, params...)
}

//TaxiReceiptRecognize 出租车票识别
func (oc *OCRClient) TaxiReceiptRecognize(image interface{}, params ...RequestParam) (*OCRResponse, error) {
	return oc.ocr(image, OCR_TAX_RECEIPT_URL, defaultTaxiReceiptParams, params...)
}

//DrivingLicenseRecognize 驾驶证识别
//识别机动车驾驶证所有关键字段
func (oc *OCRClient) DrivingLicenseRecognize(image interface{}, params ...RequestParam) (*OCRResponse, error) {
	return oc.ocr(image, OCR_DRIVING_LICENSE_URL, defaultDrivingLicenseParams, params...)
}

//VehicleLicenseRecognize 行驶证识别
//识别机动车行驶证所有关键字段
func (oc *OCRClient) VehicleLicenseRecognize(image interface{}, params ...RequestParam) (*OCRResponse, error) {
	return oc.ocr(image, OCR_VEHICLE_LICENSE_URL, defaultVehicleLicenseParams, params...)
}

//LicensePlateRecognize 车牌识别
//对小客车的车牌进行识别
func (oc *OCRClient) LicensePlateRecognize(image interface{}, params ...RequestParam) (*OCRResponse, error) {
	return oc.ocr(image, OCR_LICENSE_PLATE_URL, defaultLicensePlateParams, params...)
}

//VinCodeRecognize VIN码识别
func (oc *OCRClient) VinCodeRecognize(image interface{}, params ...RequestParam) (*OCRResponse, error) {
	return oc.ocr(image, OCR_VIN_CODE_URL, defaultVinCodeParams, params...)
}

//NumbersRecognize 数字识别
func (oc *OCRClient) NumbersRecognize(image interface{}, params ...RequestParam) (*OCRResponse, error) {
	return oc.ocr(image, OCR_NUMBERS_URL, defaultNumbersParams, params...)
}

//WebimageRecognize 网络图片识别
//识别一些网络上背景复杂，特殊字体的文字
func (oc *OCRClient) WebimageRecognize(image interface{}, params ...RequestParam) (*OCRResponse, error) {
	return oc.ocr(image, OCR_WEBIMAGE_URL, defaultWebimgParams, params...)
}

////GeneralRecognizeWithLocation 通用文字识别（含位置信息）
////识别图片中的文字信息（包含文字区域的坐标信息）
//func (oc *OCRClient) GeneralRecognizeWithLocation(image interface{}, params ...RequestParam) (*OCRResponse, error) {
//
//	return oc.ocr(image, OCR_GENERAL_WITH_LOCATION_URL, defaultGeneralWithLocationParams, params...)
//
//}
//
////GeneralRecognizeEnhanced 通用文字识别（含生僻字）
////识别图片中的文字信息（包含对常见字和生僻字的识别）
//func (oc *OCRClient) GeneralRecognizeEnhanced(image interface{}, params ...RequestParam) (*OCRResponse, error) {
//
//	return oc.ocr(image, OCR_GENERAL_ENHANCED_URL, defaultDeneralEnhancedParams, params...)
//
//}

func (oc *OCRClient) ocr(image interface{}, url string, def map[string]interface{}, params ...RequestParam) (*OCRResponse, error) {
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
