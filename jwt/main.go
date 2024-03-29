package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var SECRET = []byte("secret-auth-key")
var api_key = "1234"

func SayHello(w http.ResponseWriter, r *http.Request) {
	//w.Write([]byte("Hello Caller from secure Area"))
	fmt.Fprintf(w, "Hello Caller from secure Area")
}

func CreateJWT() (string, error) {

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["exp"] = time.Now().Add(time.Hour).Unix()
	claims["username"] = "Dushyant"

	tokenStr, err := token.SignedString(SECRET)

	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}

	return tokenStr, nil
}

func ValidateJWT(next func(w http.ResponseWriter, r *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header["Token"] != nil {
			token, err := jwt.Parse(r.Header["Token"][0], func(t *jwt.Token) (interface{}, error) {
				_, ok := t.Method.(*jwt.SigningMethodHMAC)
				if !ok {
					w.WriteHeader(http.StatusUnauthorized)
					w.Write([]byte("not authorized"))
				}
				return SECRET, nil
			})

			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("not authorized\n"))
			}

			if token.Valid {
				next(w, r)
			}
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("not authorized\n"))
		}
	})
}

func GetJWT(w http.ResponseWriter, r *http.Request) {
	if r.Header["Access"] != nil {
		if r.Header["Access"][0] != api_key {
			return
		} else {
			jwtStr, err := CreateJWT()
			if err != nil {
				return
			}
			fmt.Fprint(w, jwtStr)
		}
	}
}

func main() {
	http.Handle("/hello", ValidateJWT(SayHello))
	http.HandleFunc("/getJwt", GetJWT)

	http.ListenAndServe(":8080", nil)
}
