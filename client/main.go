package main

import "fmt"

import jwt "github.com/dgrijalva/jwt-go"

var secret = []byte("topsectethjbdvkjdsdn67478")

func generate() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claim := token.Claims.(jwt.MapClaims)
	claim["name"] = "akash"
	claim["age"] = 19
	claim["authorised"] = true
	tokenstring, err := token.SignedString(secret)
	if err != nil {
		fmt.Errorf("error %s", err.Error())
		return "", nil
	}
	return tokenstring, nil
}
func main() {
	fmt.Println("my simple client")
	tokenstring, err := generate()
	if err != nil {
		fmt.Println("error generating token")
	}
	fmt.Printf("%s", tokenstring)
}
