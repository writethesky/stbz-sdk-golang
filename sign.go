package stbz

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

// Sign 签名
func Sign(params signV2, queryMaps map[string]string, bodyMaps map[string]interface{}, secretKey string) string {
	bodyBytes, err := json.Marshal(bodyMaps)
	if nil != err {
		return ""
	}
	body := string(bodyBytes)
	// if len(bodyMaps) == 0 {
	// 	body = ""
	// }

	return createSignV2(params, queryMaps, body, secretKey)
}

type signV2 struct {
	AppKey    string `header:"Api-App-Key"`
	TimeStamp uint64 `header:"Api-Time-Stamp"`
	Nonce     string `header:"Api-Nonce"`
	Sign      string `header:"Api-Sign"`
}

func (s signV2) toMap() (maps map[string]string) {
	maps = make(map[string]string)
	rtype := reflect.TypeOf(s)
	rvalue := reflect.ValueOf(&s)
	for i := 0; i < rtype.NumField(); i++ {
		field := rtype.Field(i)
		// field.Type.Kind()

		switch rvalue.Elem().Field(i).Interface().(type) {
		case string:
			maps[field.Tag.Get("header")] = rvalue.Elem().Field(i).Interface().(string)
		case uint64:
			maps[field.Tag.Get("header")] = strconv.FormatUint(rvalue.Elem().Field(i).Interface().(uint64), 10)
		}

	}
	return
}

func (s signV2) toParamsMap() (maps map[string]string) {
	maps = make(map[string]string)
	rtype := reflect.TypeOf(s)
	rvalue := reflect.ValueOf(&s)
	for i := 0; i < rtype.NumField(); i++ {
		field := rtype.Field(i)
		// field.Type.Kind()
		if field.Name == "Sign" {
			continue
		}
		switch rvalue.Elem().Field(i).Interface().(type) {
		case string:
			maps[field.Tag.Get("header")] = rvalue.Elem().Field(i).Interface().(string)
		case uint64:
			maps[field.Tag.Get("header")] = strconv.FormatUint(rvalue.Elem().Field(i).Interface().(uint64), 10)
		}

	}
	return
}

func getSignv2Params(header http.Header) (params signV2, err error) {

	rtype := reflect.TypeOf(params)
	rvalue := reflect.ValueOf(&params)
	for i := 0; i < rtype.NumField(); i++ {

		field := rtype.Field(i)
		headerTag := field.Tag.Get("header")
		headerValue := header.Get(headerTag)

		if "" == headerValue {
			err = errors.New("必须传递 " + headerTag)
			return
		}
		switch field.Type.Kind().String() {
		case "string":

			rvalue.Elem().Field(i).Set(reflect.ValueOf(headerValue))

		case "uint64":
			headerIntValue, err2 := strconv.ParseUint(headerValue, 10, 64)
			if nil != err2 {
				fmt.Println(err2)
				err = errors.New("必须传递 合法的 " + headerTag)
				return
			}

			rvalue.Elem().Field(i).Set(reflect.ValueOf(headerIntValue))

		}

	}

	// queryParams := c.QueryParams()
	// for k, _ := range queryParams {
	// 	params[k] = queryParams.Get(k)

	// }

	// params["Api-App-Key"] = appKey
	// params["Api-Time-Stamp"] = timestamp
	// params["Api-Nonce"] = nonce
	// params["Api-Sign"] = sign

	// timestampInt, err := strconv.Atoi(timestamp)
	// if nil != err {
	// 	code = code.UnKnow()
	// 	return
	// }
	// cstSh, _ := time.LoadLocation("Asia/Shanghai")
	// currentTime := int(time.Now().In(cstSh).UnixNano() / 1e6)
	// if nil != err || timestampInt > currentTime || timestampInt+600000 < currentTime {
	// 	// code = code.SignErr("签名过期或时间还未到")
	// 	// return
	// }

	// code = code.Success()
	return
}

// 生成签名
func createSignV2(params signV2, queryMaps map[string]string, body string, secretKey string) (sign string) {
	allMaps := make(map[string]string)

	for k, v := range queryMaps {
		allMaps[k] = v
	}

	for k, v := range params.toParamsMap() {
		allMaps[k] = v
	}

	keys := make([]string, 0)
	for k := range allMaps {
		keys = append(keys, k)

	}

	sort.Strings(keys)

	paramsString := ""
	for _, k := range keys {
		paramsString += k + allMaps[k]
	}

	re3, _ := regexp.Compile(`\s`)
	body = re3.ReplaceAllString(body, "")

	paramsString += secretKey + body

	sha1 := sha1.New()
	sha1.Write([]byte(paramsString))
	sha1String := hex.EncodeToString(sha1.Sum([]byte("")))

	md5 := md5.New()
	md5.Write([]byte(sha1String))
	md5String := hex.EncodeToString(md5.Sum([]byte("")))
	sign = strings.ToUpper(md5String)

	// fmt.Println("参与签名计算的参数：" + paramsString)
	// fmt.Println("sha1结果：" + sha1String)
	// fmt.Println("md5结果：" + md5String)
	// fmt.Println("签名结果：" + sign)

	return
}
