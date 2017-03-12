package handler

import(
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"
	"crypto/rsa"

	"github.com/dgrijalva/jwt-go"
	
	"github.com/wuriyanto48/go-oauth2-jwt/response"
	
)

type UserLogin struct{
	Username string `json:"username"`
	Password string `json:"password"`
}

type AccessToken struct{
	AccessToken string `json:"access_token"`
}

func GetAccessToken(signKey *rsa.PrivateKey) func(res http.ResponseWriter, req *http.Request){
		
		return func(res http.ResponseWriter, req *http.Request){
			
			queries := req.URL.Query()
			query, _ := queries["grant_type"]
			grantType := query[0]
			
			switch(grantType){
				case "password":
					var userLogin UserLogin
					decoder := json.NewDecoder(req.Body)
					err := decoder.Decode(&userLogin)
					if err != nil {
						response.MessageWithJson(res, "Error occured", http.StatusInternalServerError)
						log.Fatal(err)
						return
					}
					if strings.ToLower(userLogin.Username) != "wuriyanto"{
						if userLogin.Password != "123456"{
							response.MessageWithJson(res, "Username or Password invalid", http.StatusUnauthorized)
							return
						}
					} else {
						token := jwt.New(jwt.SigningMethodRS256)
						claims := make(jwt.MapClaims)
						claims["iss"] = "auth.wury.com"
						claims["aud"] = "wury.com"
						claims["exp"] = time.Now().Add(time.Minute * 10).Unix()
						claims["sub"] = "001"
						token.Claims = claims
						
						tokenString, err := token.SignedString(signKey)
						if err != nil {
							response.MessageWithJson(res, "Can't get access token", http.StatusInternalServerError)
							log.Fatal(err)
							return
						}
						
						response.JsonResponse(res, AccessToken{tokenString}, http.StatusOK)
					}
					
					
				default:
					response.MessageWithJson(res, "Invalid grant type", http.StatusNotFound)
					return
			}
		}
}




