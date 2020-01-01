package utils

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"lemon/utils/crypto"
	"net/url"
	"sort"
)

// 快速打印出一个变量，直接退出，加速调试
func PanicJson(a interface{}) {
	bs, err := json.MarshalIndent(a, "", "\t")
	if err != nil {
		panic(err)
	}
	panic(string(bs))
}

// 生成签名
func CreateSign(params url.Values) string {
	var key []string
	var str = ""
	for k := range params {
		if k != "sign" {
			key = append(key, k)
		}
	}
	sort.Strings(key)

	for i := 0; i < len(key); i++ {
		if i == 0 {
			str = fmt.Sprintf("%v=%v", key[i], params.Get(key[i]))
		} else {
			str = str + fmt.Sprintf("&%v=%v", key[i], params.Get(key[i]))
		}
	}

	// 自定义签名算法
	return crypto.MD5(crypto.MD5(str) + crypto.MD5(viper.GetString("default_db.addr")))
}
