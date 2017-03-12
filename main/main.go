package main

import(
	"log"
	"net/http"
	
	"github.com/gorilla/mux"
	"github.com/wuriyanto48/go-oauth2-jwt/handler"
)

func main(){
	
	server()
}

func server(){
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", handler.Index).Methods("GET")
	log.Fatal(http.ListenAndServe(":3000", router))
}