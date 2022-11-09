package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type MyClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func main() {
	mySigningKey := []byte("woshizccc")
	c := MyClaims{
		Username: "zc",
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 60, //一分钟前生效
			ExpiresAt: time.Now().Unix() + 60*60*2,
			Issuer:    "zc",
		},
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	fmt.Println(t)
	s, e := t.SignedString(mySigningKey) //签发
	if e != nil {
		fmt.Printf("%s", e)
	}
	fmt.Println(s)
	token, err := jwt.ParseWithClaims(s, &MyClaims{}, func(token *jwt.Token) (interface{}, error) { //解密
		return mySigningKey, nil
	})
	fmt.Println(err)
	fmt.Println(token.Claims.(*MyClaims)) //断言成指针

}
