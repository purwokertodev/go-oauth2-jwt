package main

import (
	"fmt"
	"strconv"
	"crypto/rsa"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/wuriyanto48/go-oauth2-jwt/handler"
	"github.com/wuriyanto48/go-oauth2-jwt/config"
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
	conv, err := config.GetConfig()
	if err != nil {
		fmt.Println("Cannot get config ", err)
	}
	port := strconv.Itoa(conv.Dev.Port)
	host := conv.Dev.Host
	fmt.Println("app running on "+host+":"+port)
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", handler.Index).Methods("GET")
	router.HandleFunc("/token", handler.GetAccessToken(signKey)).Methods("POST")
	log.Fatal(http.ListenAndServe(host+":"+port, router))
}
