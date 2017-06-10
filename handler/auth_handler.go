package handler

import (
	"crypto/rsa"
	"encoding/json"
	"os"
	"log"
	"net/http"
	"time"

	"github.com/wuriyanto48/go-json-message/response"
	"github.com/wuriyanto48/go-oauth2-jwt/login"
	"github.com/wuriyanto48/go-oauth2-jwt/token"
)

func GetAccessToken(signKey *rsa.PrivateKey) func(res http.ResponseWriter, req *http.Request) {

	return func(res http.ResponseWriter, req *http.Request) {

		queries := req.URL.Query()
		query, _ := queries["grant_type"]
		grantType := query[0]

		switch grantType {
		case "password":
			var userLogin login.UserLogin
			decoder := json.NewDecoder(req.Body)
			err := decoder.Decode(&userLogin)
			if err != nil {
				response.MessageWithJson(res, "Error occured", http.StatusInternalServerError)
				log.Fatal(err)
				return
			}
			if !userLogin.IsValidUser() {
				response.MessageWithJson(res, "Username or Password invalid", http.StatusUnauthorized)
				return
			} else {
				tokenExpired, err := time.ParseDuration(os.Getenv("TOKEN_EXPIRED_MINUTES"))
				if err != nil {
					panic("cant parse time duration")
				}
				claim := token.NewClaim("auth.wury.com", "wury.com", "001", tokenExpired)
				tokenResult := <-claim.GenerateToken(signKey)
				if tokenResult.Error != nil {
					response.MessageWithJson(res, "Cant get token", http.StatusInternalServerError)
					return
				}
				response.JsonResponse(res, tokenResult.AccessToken, http.StatusOK)
			}

		default:
			response.MessageWithJson(res, "Invalid grant type", http.StatusNotFound)
			return
		}
	}
}
