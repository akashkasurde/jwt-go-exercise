package main

import (
	"fmt"
	"log"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
)

var signingkey = []byte("topsectethjbdvkjdsdn67478")

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, string([]byte("super secret information")))
}

func isallowed(handler func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header["Token"] != nil {
			token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("something gone wrong")
				}
				return signingkey, nil
			})

			if err != nil {
				fmt.Errorf("wrong gone something")
			}
			if token.Valid {
				handler(w, r)
			}
		} else {
			fmt.Fprintf(w, "not allowed to access")
		}
	})
}

func handlerequest() {
	http.HandleFunc("/", isallowed(homepage))
	log.Fatal(http.ListenAndServe(":9001", nil))
}
func main() {
	fmt.Printf("go to :9001")
	handlerequest()
}
