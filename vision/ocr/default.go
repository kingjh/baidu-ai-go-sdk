package ocr

var defaultGeneralBasicParams = map[string]interface{}{
	"image":            "",        //图像数据，base64编码，要求base64编码后大小不超过4M，最短边至少15px，最长边最大4096px,支持jpg/png/bmp格式
	"url":              "",        //图片完整URL，URL长度不超过1024字节，URL对应的图片base64编码后大小不超过4M，最短边至少15px，最长边最大4096px,支持jpg/png/bmp格式，当image字段存在时url字段失效，不支持https的图片链接
	"language_type":    "CHN_ENG", //识别语言类型，默认为CHN_ENG。可选值包括： - CHN_ENG：中英文混合； - ENG：英文； - POR：葡萄牙语； - FRE：法语； - GER：德语； - ITA：意大利语； - SPA：西班牙语； - RUS：俄语； - JAP：日语； - KOR：韩语
	"detect_direction": "false",   //是否检测图像朝向，默认不检测，即：false。朝向是指输入图像是正常方向、逆时针旋转90/180/270度。可选值包括: - true：检测朝向； - false：不检测朝向。
	"detect_language":  "false",   //是否检测语言，默认不检测。当前支持（中文、英语、日语、韩语）
	"probability":      "false",   //是否返回识别结果中每一行的置信度
}

var defaultHandwritingParams = map[string]interface{}{
	"image":                 "", //图像数据，base64编码，要求base64编码后大小不超过4M，最短边至少15px，最长边最大4096px,支持jpg/png/bmp格式
	"recognize_granularity": "", //是否定位单字符位置，big：不定位单字符位置，默认值；small：定位单字符位置
	"words_type":            "", //words_type=number:手写数字识别；无此参数或传其它值 默认手写通用识别（目前支持汉字和英文）
}

var defaultIdcardParams = map[string]interface{}{
	"image":            "",      //图像数据，base64编码，要求base64编码后大小不超过4M，最短边至少15px，最长边最大4096px,支持jpg/png/bmp格式
	"detect_direction": "false", //是否检测图像朝向，默认不检测，即：false。朝向是指输入图像是正常方向、逆时针旋转90/180/270度。可选值包括: - true：检测朝向； - false：不检测朝向。
	"id_card_side":     "front", //front：身份证正面；back：身份证背面
	"detect_risk":      "false", //是否开启身份证风险类型(身份证复印件、临时身份证、身份证翻拍、修改过的身份证)功能，默认不开启，即：false。可选值:true-开启；false-不开启
}

var defaultBankcardParams = map[string]interface{}{
	"image": "", //图像数据，base64编码，要求base64编码后大小不超过4M，最短边至少15px，最长边最大4096px,支持jpg/png/bmp格式
}

var defaultBusinessLicenseParams = map[string]interface{}{
	"image":            "",      //图像数据，base64编码，要求base64编码后大小不超过4M，最短边至少15px，最长边最大4096px,支持jpg/png/bmp格式
	"detect_direction": "false", //是否检测图像朝向，默认不检测，即：false。朝向是指输入图像是正常方向、逆时针旋转90/180/270度。可选值包括: - true：检测朝向； - false：不检测朝向。
	"accuracy":         "",      //可选值：normal,high参数选normal或为空使用快速服务；选择high使用高精度服务，但是时延会根据具体图片有相应的增加
}

var defaultPassportParams = defaultBankcardParams

var defaultFormParams = defaultBankcardParams

var defaultFormResultParams = map[string]interface{}{
	"request_id":  "",      //发送表格文字识别请求时返回的request id
	"result_type": "excel", //期望获取结果的类型，取值为“excel”时返回xls文件的地址，取值为“json”时返回json格式的字符串,默认为”excel”
}

var defaultReceiptParams = map[string]interface{}{
	"image":                 "",      //图像数据，base64编码，要求base64编码后大小不超过4M，最短边至少15px，最长边最大4096px,支持jpg/png/bmp格式
	"recognize_granularity": "",      //是否定位单字符位置，big：不定位单字符位置，默认值；small：定位单字符位置
	"probability":           "false", //是否返回识别结果中每一行的置信度
	"accuracy":              "",      //可选值：normal,high参数选normal或为空使用快速服务；选择high使用高精度服务，但是时延会根据具体图片有相应的增加
	"detect_direction":      "false", //是否检测图像朝向，默认不检测，即：false。朝向是指输入图像是正常方向、逆时针旋转90/180/270度。可选值包括: - true：检测朝向； - false：不检测朝向。
}

var defaultVatInvoiceParams = map[string]interface{}{
	"image":    "", //图像数据，base64编码，要求base64编码后大小不超过4M，最短边至少15px，最长边最大4096px,支持jpg/png/bmp格式
	"accuracy": "", //可选值：normal,high参数选normal或为空使用快速服务；选择high使用高精度服务，但是时延会根据具体图片有相应的增加
}

var defaultTrainTicketParams = defaultBankcardParams

var defaultTaxiReceiptParams = defaultBankcardParams

var defaultDrivingLicenseParams = map[string]interface{}{
	"image":                "",      //图像数据，base64编码，要求base64编码后大小不超过4M，最短边至少15px，最长边最大4096px,支持jpg/png/bmp格式
	"detect_direction":     "false", //是否检测图像朝向，默认不检测，即：false。朝向是指输入图像是正常方向、逆时针旋转90/180/270度。可选值包括: - true：检测朝向； - false：不检测朝向。
	"unified_valid_period": "true",  //true: 归一化格式输出；false 或无此参数按非归一化格式输出
}

var defaultVehicleLicenseParams = map[string]interface{}{
	"image":            "",      //图像数据，base64编码，要求base64编码后大小不超过4M，最短边至少15px，最长边最大4096px,支持jpg/png/bmp格式
	"detect_direction": "false", //是否检测图像朝向，默认不检测，即：false。朝向是指输入图像是正常方向、逆时针旋转90/180/270度。可选值包括: - true：检测朝向； - false：不检测朝向。
	"accuracy":         "",      //normal 使用快速服务，1200ms左右时延；缺省或其它值使用高精度服务，1600ms左右时延
}

var defaultLicensePlateParams = map[string]interface{}{
	"image":        "", //图像数据，base64编码，要求base64编码后大小不超过4M，最短边至少15px，最长边最大4096px,支持jpg/png/bmp格式
	"multi_detect": "", //是否检测多张车牌，默认为false，当置为true的时候可以对一张图片内的多张车牌进行识别
}

var defaultVinCodeParams = map[string]interface{}{
	"image":        "",     //图像数据，base64编码，要求base64编码后大小不超过4M，最短边至少15px，最长边最大4096px,支持jpg/png/bmp格式
	"imgDirection": "true", //该参数决定模型是否自动纠正图像方向，默认不检测，当该参数=setImgDirFlag 时自动检测、纠正图像方向并识别
}

var defaultNumbersParams = map[string]interface{}{
	"image":                 "",      //图像数据，base64编码，要求base64编码后大小不超过4M，最短边至少15px，最长边最大4096px,支持jpg/png/bmp格式
	"recognize_granularity": "",      //是否定位单字符位置，big：不定位单字符位置，默认值；small：定位单字符位置
	"detect_direction":      "false", //是否检测图像朝向，默认不检测，即：false。朝向是指输入图像是正常方向、逆时针旋转90/180/270度。可选值包括: - true：检测朝向； - false：不检测朝向。
}

var defaultWebimgParams = map[string]interface{}{
	"image":            "",      //图像数据，base64编码，要求base64编码后大小不超过4M，最短边至少15px，最长边最大4096px,支持jpg/png/bmp格式
	"detect_direction": "false", //是否检测图像朝向，默认不检测，即：false。朝向是指输入图像是正常方向、逆时针旋转90/180/270度。可选值包括:- true：检测朝向； - false：不检测朝向
	"detect_language":  "false", //是否检测语言，默认不检测。当前支持（中文、英语、日语、韩语）
}

//var defaultGeneralWithLocationParams = map[string]interface{}{
//	"image":                 "",        //图像数据，base64编码，要求base64编码后大小不超过4M，最短边至少15px，最长边最大4096px,支持jpg/png/bmp格式
//	"url":                   "",        //图片完整URL，URL长度不超过1024字节，URL对应的图片base64编码后大小不超过4M，最短边至少15px，最长边最大4096px,支持jpg/png/bmp格式，当image字段存在时url字段失效，不支持https的图片链接
//	"recognize_granularity": "big",     //是否定位单字符位置，big：不定位单字符位置，默认值；small：定位单字符位置
//	"language_type":         "CHN_ENG", //识别语言类型，默认为CHN_ENG。可选值包括： - CHN_ENG：中英文混合； - ENG：英文； - POR：葡萄牙语； - FRE：法语； - GER：德语； - ITA：意大利语； - SPA：西班牙语； - RUS：俄语； - JAP：日语
//	"detect_direction":      "false",   //是否检测图像朝向，默认不检测，即：false。朝向是指输入图像是正常方向、逆时针旋转90/180/270度。可选值包括:- true：检测朝向； - false：不检测朝向
//	"detect_language":       "false",   //是否检测语言，默认不检测。当前支持（中文、英语、日语、韩语）
//	"vertexes_location":     "false",   //是否返回文字外接多边形顶点位置，不支持单字位置。默认为false
//	"probability":           "false",   //是否返回识别结果中每一行的置信度
//}
//
//var defaultDeneralEnhancedParams = map[string]interface{}{
//	"image":            "",        //图像数据，base64编码，要求base64编码后大小不超过4M，最短边至少15px，最长边最大4096px,支持jpg/png/bmp格式
//	"language_type":    "CHN_ENG", //识别语言类型，默认为CHN_ENG。可选值包括： - CHN_ENG：中英文混合； - ENG：英文； - POR：葡萄牙语； - FRE：法语； - GER：德语； - ITA：意大利语； - SPA：西班牙语； - RUS：俄语； - JAP：日语
//	"detect_direction": "false",   //是否检测图像朝向，默认不检测，即：false。朝向是指输入图像是正常方向、逆时针旋转90/180/270度。可选值包括:- true：检测朝向； - false：不检测朝向
//	"detect_language":  "false",   //是否检测语言，默认不检测。当前支持（中文、英语、日语、韩语）
//	"probability":      "false",   //是否返回识别结果中每一行的置信度
//}
//
