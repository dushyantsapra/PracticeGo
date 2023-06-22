package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func SayHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Caller"))
}

var SECRET = []byte("super-secret-auth-key")
var api_key = "1234"

func CreateJWT() (string, error) {

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["exp"] = time.Now().Add(time.Hour).Unix()

	tokenStr, err := token.SignedString(SECRET)

	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}

	return tokenStr, nil
}

func main() {
	http.HandleFunc("/hello", SayHello)

	http.ListenAndServe(":8080", nil)
}
