package main

import (
	"crypto/rsa"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/wuriyanto48/go-oauth2-jwt/handler"
)

const (
	privateKeyPath = "secret/app.rsa"
	publicKeyPath  = "secret/app.rsa.pub"
)

var (
	verifyKey *rsa.PublicKey
	signKey   *rsa.PrivateKey
)

func main() {
	initKeys()
	server()
}

func initKeys() {
	signBytes, err := ioutil.ReadFile(privateKeyPath)
	printError(err)

	signKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	printError(err)

	verifyBytes, err := ioutil.ReadFile(publicKeyPath)
	printError(err)

	verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	printError(err)
}

func printError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func server() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", handler.Index).Methods("GET")
	router.HandleFunc("/token", handler.GetAccessToken(signKey)).Methods("POST")
	log.Fatal(http.ListenAndServe(":3000", router))
}
