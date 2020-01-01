package utils

import (
	"fmt"
	"net/url"
	"testing"
)

func TestCreateSign(t *testing.T) {
	params := url.Values{
		"method": {"get"},
		"id":     {"1"},
	}
	fmt.Println(CreateSign(params))
}
