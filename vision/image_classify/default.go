package image_classify

var defaultDishParams = map[string]interface{}{
	"top_num":          "5",    //返回结果top n,默认5.
	"filter_threshold": "0.95", //默认0.95，可以通过该参数调节识别效果，降低非菜识别率.
}
