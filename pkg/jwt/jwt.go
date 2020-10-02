package jwt

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"strconv"
	"time"
)

type CustomClaims struct {
	Uid int64 `json:"uid"`
	jwt.StandardClaims
}

func CreateToken(uid int64) (string, error) {
	string_uid := strconv.FormatInt(uid, 10)
	claims := CustomClaims{
		uid,
		jwt.StandardClaims{
			Id:        string_uid,
			Subject:   "MySelfBlog",
			Audience:  "MySelfBlog",
			ExpiresAt: time.Now().Unix() + (24 * 3600 * 7),
			Issuer:    "",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString([]byte("MySelfBlog"))
	//tokenStr, err := token.SignedString([]byte("946356eb-204d-4be3-8eb0-32720a403814"))
	if err != nil {
		return "", errors.New(fmt.Sprintf(`create token err:%v`, err))
	}
	return tokenStr, err
}

func VerifyToken(token_string string) (uid int64, err error) {
	tokenValue, err := jwt.ParseWithClaims(token_string, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(`MySelfBlog`), nil
	})
	if err != nil {
		err = errors.New(err.Error())
		return
	}
	claims, ok := tokenValue.Claims.(*CustomClaims)
	if !ok {
		err = errors.New("Token is invalid")
		return
	}
	uid = claims.Uid
	return
}
