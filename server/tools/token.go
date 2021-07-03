package tools

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

/*
	jwt.SigningMethodHS256 和 jwt.SigningMethodHMAC 一对
*/

var (
	secretKey = ""
)

func newSecretKey() []byte {
	if len(os.Getenv("Security")) == 0 {
		secretKey = "xzc45sv1r3&2xd@sc&*14a~~~sdzRZXv"
	}
	return []byte(secretKey)
}

func NewToken(v interface{}) (string, bool) {
	jsonStr, err := json.Marshal(v)
	if err != nil {
		return "", false
	}

	claims := jwt.MapClaims{
		"exp":  time.Now().Add(time.Hour * 72).Unix(),
		"iat":  time.Now().Unix(),
		"data": string(jsonStr),
	}

	//                                       算法            负载payload               盐值
	tokenString, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(newSecretKey())
	if err != nil {
		return "", false
	}

	return tokenString, true
}
























func CheckToken(tokenString string) (string, bool) {
	// 检测算法,返回盐值解析得到token struct(string)对象
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (i interface{}, e error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return newSecretKey(), nil
	})
	if err != nil {
		return "", false
	}

	// 解析payload类型				MapClaims => map[string]interface
	if claims, ok := token.Claims.(jwt.MapClaims); ok && int64(claims["exp"].(float64)) >= time.Now().Unix() {
		return claims["data"].(string), true
	}
	return "", false
}

func GetTokenMapClaims(tokenString string) map[string]interface{} {
	// 检测算法,返回盐值解析得到token struct(string)对象
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (i interface{}, e error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return newSecretKey(), nil
	})
	if err != nil {
		return nil
	}

	// 解析payload类型				MapClaims => map[string]interface
	if claims, ok := token.Claims.(jwt.MapClaims); ok && int64(claims["exp"].(float64)) >= time.Now().Unix() {
		return claims
	}
	return nil
}

