package utils

import (
	"fmt"
	"encoding/json"
	"time"
	"crypto/md5"
	"encoding/hex"
	"net/url"
	"sort"
	"github.com/spf13/viper"
	"os"
)

// 快速打印出一个变量，直接退出，加速调试
func PanicJson(a interface{}) {
	bs, err := json.MarshalIndent(a, "", "\t")
	if err != nil {
		panic(err)
	}
	panic(string(bs))
}

// 检查文件是否存在
func FileIsExist(file string) bool {
	if _, err := os.Stat(file); os.IsExist(err) || err == nil {
		return true
	}
	return false
}

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
	return MD5(MD5(str) + MD5(viper.GetString("default_db.addr")))
}
