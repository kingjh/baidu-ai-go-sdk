package image_classify

import "strconv"

type RequestParam func(map[string]interface{})

//返回结果top n,默认5
func TopNum(n uint32) RequestParam {
	return func(m map[string]interface{}) {
		m["top_num"] = strconv.Itoa(int(n))
	}
}

//默认0.95，可以通过该参数调节识别效果，降低非菜识别率.
func FilterThreshold(f float64) RequestParam {
	return func(m map[string]interface{}) {
		m["filter_threshold"] = strconv.FormatFloat(float64(f), 'f', 6, 64)
	}
}

//返回百科信息的结果数，默认不返回
func BaikeNum(n uint32) RequestParam {
	return func(m map[string]interface{}) {
		m["baike_num"] = strconv.Itoa(int(n))
	}
}
