package utils

import (
	"fmt"
	"time"
	"crypto/md5"
	"encoding/hex"
	"net/url"
	"sort"
	"github.com/spf13/viper"
)

// 获取当前标准时间
func GetTimeStandar() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// 获取当前时间戳
func GetTimeUnix() int64 {
	return time.Now().Unix()
}

// MD5 方法
func MD5(str string) string {
	s := md5.New()
	s.Write([]byte(str))
	return hex.EncodeToString(s.Sum(nil))
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
	return MD5(MD5(str) + MD5(viper.GetString("local_db.addr")))
}
