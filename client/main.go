package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
)

var secret = []byte("topsectethjbdvkjdsdn67478")

func HomePage(w http.ResponseWriter, r *http.Request) {
	validtoken, err := generate()
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://localhost:9001/", nil)
	req.Header.Set("Token", validtoken)
	res, err := client.Do(req)
	if err != nil {
		fmt.Errorf("response  failed")
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("body failed")
	}
	fmt.Fprintf(w, string(body))
}
func handlerequest() {
	http.HandleFunc("/", HomePage)
	log.Fatal(http.ListenAndServe(":9000", nil))
}
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
	fmt.Printf("server listening on localhost:9000")
	handlerequest()
}
