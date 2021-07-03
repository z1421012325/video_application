package tools

import (
	"crypto/md5"

	"video_application/server/conf"
)

var salf = conf.Config.Mixture

// 密码加密
func EnCryptionPassword(password string) (crypStr string) {

	password = password + salf
	m := md5.New()
	crypStr = string(m.Sum([]byte(password)))
	return
}

// 密码检测
func CheckPassword(encryptionStr, VerifyPassword string) bool {
	return EnCryptionPassword(VerifyPassword) == encryptionStr
}




