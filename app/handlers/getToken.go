package handlers

import (
	"crypto/rsa"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	privateKeyPath = "app.rsa" // openssl genrsa -out app.rsa 2048
	access
)

var (
	signKey *rsa.PrivateKey
)

func init() {
	fmt.Println("init called.")

	signBytes, _ := ioutil.ReadFile(privateKeyPath)
	signKey, _ = jwt.ParseRSAPrivateKeyFromPEM(signBytes)

}

// GetTokenHandler is a handler function for '/get-token' route
func GetTokenHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Method != "POST" {
			http.Error(w, "No POST", http.StatusBadRequest)
			return
		}

		token := jwt.New(jwt.GetSigningMethod("RS256"))

		claims := make(jwt.MapClaims)
		claims["AccessToken"] = "admin"
		claims["name"] = "ggg"
		claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
		claims["iat"] = time.Now().Unix()
		token.Claims = claims

		tokenString, err := token.SignedString(signKey)

		if err != nil {
			http.Error(w, "No POST", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(tokenString))

	})
}
