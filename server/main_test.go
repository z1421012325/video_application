package main

import (
	"crypto/md5"
	"testing"
	"video_application/server/tools"
)

// 测试密码加密
func TestPassword(t *testing.T) {
	mixture := "999"
	passwordStr := "123465789"

	EnCryptionPassword := func () string {
		m := md5.New()
		crypStr := string(m.Sum([]byte(passwordStr+mixture)))
		return crypStr
	}

	if !tools.CheckPassword(EnCryptionPassword(),passwordStr+mixture){
		t.Fatal("密码验证错误")
	}
	t.Log("success!")
}