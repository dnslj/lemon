package crypto

import "golang.org/x/crypto/bcrypt"

// 使用bcrypt加密字符串
func Encrypt(source string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(source), bcrypt.DefaultCost)
	return string(hashedBytes), err
}

// 对比原字符串和加密后的字符串是否匹配
func Compare(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
