package util

import (
	"douyin-backend/config"
	"douyin-backend/entity"
	"github.com/dgrijalva/jwt-go"
	"log"
	"strconv"
	"time"
)

// NewToken 将userPO信息放入token的对应区域
func NewToken(userPO entity.UserPO) string {
	expiresTime := time.Now().Unix() + int64(config.ONEDAY)
	id64 := userPO.Id
	claims := jwt.StandardClaims{
		Audience:  userPO.Name,
		ExpiresAt: expiresTime,
		Id:        strconv.FormatInt(id64, 10),
		IssuedAt:  time.Now().Unix(),
		Issuer:    "tiktok",
		NotBefore: time.Now().Unix(),
		Subject:   "token",
	}
	log.Println("token claim:", claims)
	//秘钥加密
	var jwtSecret = []byte(config.SECRET_KEY)
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	if token, err := tokenClaims.SignedString(jwtSecret); err == nil {
		return token
	} else {
		log.Println("token生成失败")
		return "fail"
	}
}

// ParseToken 解析token 返回claim
func ParseToken(token string) (*jwt.StandardClaims, error) {
	jwtToken, err := jwt.ParseWithClaims(token, &jwt.StandardClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return []byte(config.SECRET_KEY), nil
	})
	if err == nil && jwtToken != nil {
		if claim, ok := jwtToken.Claims.(*jwt.StandardClaims); ok && jwtToken.Valid {
			log.Println("解析token成功")
			return claim, nil
		}
	}
	log.Println("解析token时出错", err.Error())
	return nil, err

}
