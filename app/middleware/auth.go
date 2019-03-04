package middleware

import (
	"crypto/rsa"
	"io/ioutil"
	"log"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

const (
	privateKeyPath = "app.rsa" // openssl genrsa -out app.rsa 2048
)

var (
	signKey *rsa.PrivateKey
)

func init() {   

	signBytes, err := ioutil.ReadFile(privateKeyPath)

	if err != nil {
		log.Fatal(err)
	}

	signKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)

	if err != nil {
		log.Fatal(err)
	}

}

func createToken(user string) (string, error) {

	token := jwt.New(jwt.GetSigningMethod("RS256"))

	claims := make(jwt.MapClaims)
	claims["user"] = user
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	claims["iat"] = time.Now().Unix()
	token.Claims = claims

	return token.SignedString(signKey)

}
