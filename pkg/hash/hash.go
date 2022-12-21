package hash

import (
	"cloud-api-go/pkg/logger"
	"golang.org/x/crypto/bcrypt"
)

// BcryptHash 对密码进行加密
func BcryptHash(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		logger.LogIf(err)
	}

	return string(bytes)
}

// BcryptCheck 对比明文密码和数据库的哈希值
func BcryptCheck(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err != nil
}

// BcryptIsHashed 检查字符串是否是哈希过
func BcryptIsHashed(str string) bool {
	return len(str) == 60
}
