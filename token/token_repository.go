package token

import (
	"crypto/rsa"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claim struct {
	Issuer   string
	Audience string
	Subject  string
}

type accessToken struct {
	AccessToken string `json:"access_token"`
}

func GenerateToken(signKey *rsa.PrivateKey, cl Claim) (interface{}, error) {
	token := jwt.New(jwt.SigningMethodRS256)
	claims := make(jwt.MapClaims)
	claims["iss"] = cl.Issuer
	claims["aud"] = cl.Audience
	claims["exp"] = time.Now().Add(time.Minute * 10).Unix()
	claims["iat"] = time.Now().Unix()
	claims["sub"] = cl.Subject
	token.Claims = claims

	tokenString, err := token.SignedString(signKey)
	if err != nil {
		return nil, err
	}
	return accessToken{tokenString}, nil
}
