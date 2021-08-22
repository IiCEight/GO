package main

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var key = []byte("i l0ve saber")

type Myclaim struct {
	Username string `json:"username`
	jwt.StandardClaims
}

func GetToken() string {
	key := []byte("i l0ve saber")
	c := Myclaim{
		Username: "Saber",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 10).Unix(),
			Issuer:    "saber",
			NotBefore: time.Now().Unix(),
		},
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	s, _ := t.SignedString(key)

	return s
}

func ParseToken(token string) (*Myclaim, error) {
	res, _ := jwt.ParseWithClaims(token, &Myclaim{}, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})

	fmt.Println(res)

	return res.Claims.(*Myclaim), nil
}

func main() {

	ParseToken(GetToken())
	return
}
