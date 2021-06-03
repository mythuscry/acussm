package common

import (
	"acussm/demo/model"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtkey =[]byte("a_secret_crect")

type claims struct {
	UserId uint
	jwt.StandardClaims
}

func ReleaseToken(user model.User)(string,error)  {

		expirationTime := time.Now().Add(7*24+time.Hour)
		claims :=&claims{
			UserId: user.ID,
			StandardClaims :jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt: time.Now().Unix(),
			Subject: "usr token",
			},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	tokenstring ,err := token.SignedString(jwtkey)
	if err != nil {
		return "", err
	}
	return tokenstring ,nil
}

func Parsetoken(tokenstring string) (*jwt.Token,*claims,error) {
	claims:=&claims{}
	token,err:=jwt.ParseWithClaims(tokenstring,claims, func(token *jwt.Token) ( i interface{}, err error) {
		return jwtkey ,nil
	})
	return token ,claims,err
}

