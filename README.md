# 胜天半子 SDK GoLang



## 使用说明

### 准备工作

1. go mod 文件增加如下配置

```golang
require github.com/writethesky/stbz-sdk-golang v1.0.0
```


2. 在需要使用sdk的地方import

```golang
import stbz "github.com/writethesky/stbz-sdk-golang"
```

3. 配置密钥
```golang
config := NewConfig("AccessKey", "SecretKey")
SetConfig(config)
```

4. 调用API
```golang
result, err := stbz.API(Method, api, queryMap, bodyMap)
```

### 使用示例

#### GET请求，query传参（form-data格式）

```golang
config := stbz.NewConfig("1AD0700D6D7D92E52ED00EDF2937BD90", "6081AE2E69B1E01DF75DA7FB77FB685E")
stbz.SetConfig(config)
result, err := stbz.API(Method.GET, "/v2/Goods/Lists", map[string]string{"page": "1", "limit": "10", "source": "0", "search_words": "牙线"}, map[string]interface{}{})

fmt.Println(result, err)

```
#### GET请求，body传参（json格式）
```golang
config := stbz.NewConfig("1AD0700D6D7D92E52ED00EDF2937BD90", "6081AE2E69B1E01DF75DA7FB77FB685E")
stbz.SetConfig(config)
result, err = stbz.API(Method.GET, "/v2/order", map[string]string{}, g.Map{"page": 1, "limit": 10, "search": g.Map{"goodsName": "蒙牛"}})
fmt.Println(result, err)

```

### POST请求，body传参（json格式）
```golang
config := stbz.NewConfig("1AD0700D6D7D92E52ED00EDF2937BD90", "6081AE2E69B1E01DF75DA7FB77FB685E")
stbz.SetConfig(config)
result, err = stbz.API(
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
fmt.Println(result, err)

```