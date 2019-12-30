package utils

import (
	"fmt"
	"net/url"
	"testing"
)

func TestFileIsExist(t *testing.T) {
	fmt.Println(FileIsExist("./utils.go"))
}

func TestGetTimeStandar(t *testing.T) {
	fmt.Println(GetTimeStandar())
}

func TestGetTimeUnix(t *testing.T) {
	fmt.Println(GetTimeUnix())
}

func TestMD5(t *testing.T) {
	fmt.Println(MD5("lemon"))
}

func TestCreateSign(t *testing.T) {
	params := url.Values{
		"method": {"get"},
		"id":     {"1"},
	}
	fmt.Println(CreateSign(params))
}
