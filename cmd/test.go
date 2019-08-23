package main

import (
	"net/url"
	"crypto/md5"
	"encoding/hex"
	"sort"
	"study/new/gindemo/config"
	"fmt"
    "math"
)

// MD5 方法
func MD5(str string) string {
	s := md5.New()
	s.Write([]byte(str))
	return hex.EncodeToString(s.Sum(nil))
}

// 生成签名
func CreateSign(params url.Values) string {
	var key []string
	for k := range params {
		if k != "sign" {
			key = append(key, k)
		}
	}

	// 排序
	sort.Strings(key)

	// 自定义签名算法
	return MD5(MD5(params.Encode()) + MD5(config.APP_NAME+config.APP_SECRET))
}

func main() {

	//v := url.Values{}
	//v.Set("z", "1")
	//v.Set("a", "1")
	//v.Set("name", "Ava")
	//v.Add("friend1", "Jess")
	//v.Add("friend2", "Sarah")
	//v.Add("friend3", "Zoe")
	//str := v.Encode()
	//fmt.Println(str)
	//fmt.Println(v.Get("name"))
	//fmt.Println(v.Get("friend"))
	//fmt.Println(v["friend"])

	//params := url.Values{
	//	"z": []string{"1"},
	//	"a": []string{"2"},
	//}
	//
	//fmt.Println(CreateSign(params))

	m := make(map[string]int, 10)
	m["a"] = 1
	m["b"] = 2
	m["sign"] = 3

	fmt.Println(m)
	delete(m, "sign")
	fmt.Println(m)
    fmt.Println(math.Float32frombits(0))
    fmt.Println(math.Float32frombits(1<<31))
    fmt.Printf("%f\n", 0.3)
    fmt.Printf("%.10f\n", 0.3)
    fmt.Printf("%.20f\n", 0.3)

}
