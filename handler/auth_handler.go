package handler

import(
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"crypto/rsa"

	"github.com/wuriyanto48/go-oauth2-jwt/response"
	"github.com/wuriyanto48/go-oauth2-jwt/token"
	
)

type UserLogin struct{
	Username string `json:"username"`
	Password string `json:"password"`
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
					if strings.ToLower(userLogin.Username) != "wuriyanto" || userLogin.Password != "123456"{
						response.MessageWithJson(res, "Username or Password invalid", http.StatusUnauthorized)
						return
					} else {
						tokenResult, err := token.GenerateToken(signKey, token.Claim{"auth.wury.com", "wury.com", "001"})
						if err != nil {
							response.MessageWithJson(res, "Cant get token", http.StatusInternalServerError)
							log.Println(err)
							return
						}
						response.JsonResponse(res, tokenResult, http.StatusOK)
					}
					
					
				default:
					response.MessageWithJson(res, "Invalid grant type", http.StatusNotFound)
					return
			}
		}
}




