package crypto

import (
	"fmt"
	"testing"
)

func TestMD5(t *testing.T) {
	fmt.Println(MD5("lemon"))
}

func TestEncrypt(t *testing.T) {
	fmt.Println(Encrypt("123"))
}

func TestCompare(t *testing.T) {
	Compare("$2a$10$t4jZ2ozfYo/9w3XhFNCsZO7d1WvwDmyGfBi..5XbelDnrySFChS5u", "123")
}
