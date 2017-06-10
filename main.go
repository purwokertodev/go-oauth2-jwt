package main

import (
	"crypto/rsa"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/wuriyanto48/go-oauth2-jwt/config"
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
	currentDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}
	configDir := flag.String("config_dir", currentDir, "find your configuration file location")

	if err := config.InitConfig(*configDir); err != nil {
		panic(err)
	}

	host := os.Getenv("HOST")
	fmt.Println("app running on " + host)
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", handler.Index).Methods("GET")
	router.HandleFunc("/token", handler.GetAccessToken(signKey)).Methods("POST")
	log.Fatal(http.ListenAndServe(host, router))
}
