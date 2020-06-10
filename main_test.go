package stbz

import (
	"testing"

	"github.com/gogf/gf/frame/g"
)

func TestMain(t *testing.T) {

	config := NewConfig("1AD0700D6D7D92E52ED00EDF2937BD90", "6081AE2E69B1E01DF75DA7FB77FB685E")
	SetConfig(config)

	// GET请求，query传参（form-data格式）的例子
	result, err := API(Method.GET, "/v2/Goods/Lists", map[string]string{"page": "1", "limit": "10", "source": "0", "search_words": "牙线"}, map[string]interface{}{})

	if nil != err || result.Code != 1 {
		t.Errorf("商品列表接口")
	}

	// GET请求，body传参（json格式）的例子
	result, err = API(Method.GET, "/v2/order", map[string]string{}, g.Map{"page": 1, "limit": 10, "search": g.Map{"goodsName": "蒙牛"}})
	if nil != err || result.Code != 1 {
		t.Errorf("订单列表接口")
	}

	// POST请求，body传参（json格式）的例子
	result, err = API(
		Method.GET,
		"/v2/order/availableCheck",
		map[string]string{},
		g.Map{
			"spu": g.Array{
				g.Map{"sku": 43870, "number": 1},
				g.Map{"sku": 429618, "number": 1},
			},
			"address": g.Map{
				"consignee":   "赵宏源",
				"phone":       "18914999333",
				"province":    "北京市",
				"city":        "北京市",
				"area":        "丰台区",
				"street":      "卢沟桥街道",
				"description": "马官营家园1号楼1单元",
			},
		},
	)
	if nil != err || result.Code != 1 {
		t.Errorf("可售性检测接口")
	}
}
